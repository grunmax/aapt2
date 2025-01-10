package main

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"io/ioutil"
	"os"
	"strings"

	"github.com/avast/apkparser"
)

func DumpBagingtOutput(manifest *Manifest) string {
	var package_data string = `package: name='%s' versionCode='%s' versionName='%s' platformBuildVersionName='%s' platformBuildVersionCode='%s' compileSdkVersion='%s' compileSdkVersionCodename='%s'
minSdkVersion:'%s'
targetSdkVersion:'%s'
application-label:'%s'
application: label='%s' icon='%s'
launchable-activity: name='%s'  label='' icon=''
`
	var launchableActivity string
	for _, activity := range manifest.Application.Activity {
		if (activity.Exported == "true") && (activity.LaunchMode == "2") {
			launchableActivity = activity.Name
			break
		}
	}

	return fmt.Sprintf(package_data, manifest.Package, manifest.VersionCode, manifest.VersionName, manifest.PlatformBuildVersionName,
		manifest.PlatformBuildVersionCode, manifest.CompileSdkVersion, manifest.CompileSdkVersionCodename, manifest.UsesSdk.MinSdkVersion,
		manifest.UsesSdk.TargetSdkVersion,
		manifest.Application.Label, manifest.Application.Label, manifest.Application.Icon,
		launchableActivity)
}

func GetManifestData(xmldata []byte) (*Manifest, error) {
	var manifest Manifest
	err := xml.Unmarshal(xmldata, &manifest)
	if err != nil {
		return nil, err
	}
	return &manifest, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Aapt2 dump badging v0.1")
		os.Exit(0)
	}
	apkName := ""
	for i := 0; i < len(os.Args); i++ {
		if strings.Contains(os.Args[i], ".apk") {
			apkName = os.Args[i]
			break
		}
	}
	if apkName == "" {
		fmt.Println("APK file is not found in parameters")
		os.Exit(0)
	}

	buf := new(bytes.Buffer)

	enc := xml.NewEncoder(buf)
	enc.Indent("", "\t")

	zipErr, resErr, manErr := apkparser.ParseApk(apkName, enc)
	if zipErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to open the APK: %s", zipErr.Error())
		os.Exit(1)
		return
	}

	if resErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse resources: %s", resErr.Error())
	}
	if manErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse AndroidManifest.xml: %s", manErr.Error())
		os.Exit(1)
		return
	}
	readBuf, _ := ioutil.ReadAll(buf)
	manifest, err := GetManifestData(readBuf)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(DumpBagingtOutput(manifest))
}
