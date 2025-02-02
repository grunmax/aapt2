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

var dump_badging_output_template string = `package: name='%s' versionCode='%s' versionName='%s' platformBuildVersionName='%s' platformBuildVersionCode='%s' compileSdkVersion='%s' compileSdkVersionCodename='%s'
minSdkVersion:'%s'
targetSdkVersion:'%s'
application-label:'%s'
application: label='%s' icon='%s'
launchable-activity: name='%s'  label='' icon=''
`

func dumpBadgingOutput(manifest *Manifest) string {
	var launchableActivity string

	// 1
	for _, activity := range manifest.Application.Activity {
		if (activity.Exported == "true") && (activity.LaunchMode == "2") {
			launchableActivity = activity.Name
			break
		}
	}
	// 2
	for _, activity := range manifest.Application.Activity {
		if (activity.Exported == "true") && (activity.LaunchMode == "1") {
			launchableActivity = activity.Name
			break
		}
	}
	// 3
	if launchableActivity == "" {
		for _, activity := range manifest.Application.Activity {
			if activity.Exported == "true" {
				launchableActivity = activity.Name
				break
			}
		}
	}
	// 4
	if launchableActivity == "" {
		for _, activity := range manifest.Application.Activity {
			if strings.Contains(activity.Name, ".Main") {
				launchableActivity = activity.Name
				break
			}
		}
	}
	// 5
	if launchableActivity == "" {
		for _, activity := range manifest.Application.Activity {
			if activity.Name != "" {
				launchableActivity = activity.Name
				break
			}
		}
	}

	return fmt.Sprintf(dump_badging_output_template, manifest.Package, manifest.VersionCode, manifest.VersionName, manifest.PlatformBuildVersionName,
		manifest.PlatformBuildVersionCode, manifest.CompileSdkVersion, manifest.CompileSdkVersionCodename, manifest.UsesSdk.MinSdkVersion,
		manifest.UsesSdk.TargetSdkVersion,
		manifest.Application.Label, manifest.Application.Label, manifest.Application.Icon,
		launchableActivity)
}

func getManifest(xmldata []byte) (*Manifest, error) {
	var manifest Manifest
	err := xml.Unmarshal(xmldata, &manifest)
	if err != nil {
		return nil, err
	}
	return &manifest, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Aapt2 dump badging v0.2")
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

	buffer := new(bytes.Buffer)

	enc := xml.NewEncoder(buffer)
	enc.Indent("", "\t")

	zipErr, resErr, manErr := apkparser.ParseApk(apkName, enc)
	if zipErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to open the APK: %s", zipErr.Error())
		os.Exit(1)
		return
	}

	if resErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse resources: %s", resErr.Error())
		os.Exit(1)
		return
	}
	if manErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse AndroidManifest.xml: %s", manErr.Error())
		os.Exit(1)
		return
	}
	data, _ := ioutil.ReadAll(buffer)
	manifest, err := getManifest(data)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(dumpBadgingOutput(manifest))
}
