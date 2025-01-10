package main

import (
	"encoding/xml"
)

type Manifest struct {
	XMLName                   xml.Name `xml:"manifest"`
	Text                      string   `xml:",chardata"`
	Android                   string   `xml:"android,attr"`
	VersionCode               string   `xml:"versionCode,attr"`
	VersionName               string   `xml:"versionName,attr"`
	CompileSdkVersion         string   `xml:"compileSdkVersion,attr"`
	CompileSdkVersionCodename string   `xml:"compileSdkVersionCodename,attr"`
	Package                   string   `xml:"package,attr"`
	PlatformBuildVersionCode  string   `xml:"platformBuildVersionCode,attr"`
	PlatformBuildVersionName  string   `xml:"platformBuildVersionName,attr"`
	UsesSdk                   struct {
		Text             string `xml:",chardata"`
		MinSdkVersion    string `xml:"minSdkVersion,attr"`
		TargetSdkVersion string `xml:"targetSdkVersion,attr"`
	} `xml:"uses-sdk"`
	UsesPermission []struct {
		Text                string `xml:",chardata"`
		Name                string `xml:"name,attr"`
		MaxSdkVersion       string `xml:"maxSdkVersion,attr"`
		UsesPermissionFlags string `xml:"usesPermissionFlags,attr"`
	} `xml:"uses-permission"`
	Queries struct {
		Text    string `xml:",chardata"`
		Package []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"package"`
		Intent []struct {
			Text   string `xml:",chardata"`
			Action struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"action"`
			Data struct {
				Text     string `xml:",chardata"`
				MimeType string `xml:"mimeType,attr"`
				Scheme   string `xml:"scheme,attr"`
				Host     string `xml:"host,attr"`
			} `xml:"data"`
			Category struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"category"`
		} `xml:"intent"`
	} `xml:"queries"`
	Permission []struct {
		Text            string `xml:",chardata"`
		Name            string `xml:"name,attr"`
		ProtectionLevel string `xml:"protectionLevel,attr"`
	} `xml:"permission"`
	UsesFeature []struct {
		Text        string `xml:",chardata"`
		GlEsVersion string `xml:"glEsVersion,attr"`
		Required    string `xml:"required,attr"`
		Name        string `xml:"name,attr"`
	} `xml:"uses-feature"`
	Application struct {
		Text                                 string `xml:",chardata"`
		Theme                                string `xml:"theme,attr"`
		Label                                string `xml:"label,attr"`
		Icon                                 string `xml:"icon,attr"`
		Name                                 string `xml:"name,attr"`
		Debuggable                           string `xml:"debuggable,attr"`
		AllowBackup                          string `xml:"allowBackup,attr"`
		HardwareAccelerated                  string `xml:"hardwareAccelerated,attr"`
		LargeHeap                            string `xml:"largeHeap,attr"`
		SupportsRtl                          string `xml:"supportsRtl,attr"`
		ExtractNativeLibs                    string `xml:"extractNativeLibs,attr"`
		ResizeableActivity                   string `xml:"resizeableActivity,attr"`
		NetworkSecurityConfig                string `xml:"networkSecurityConfig,attr"`
		AppComponentFactory                  string `xml:"appComponentFactory,attr"`
		EnableOnBackInvokedCallback          string `xml:"enableOnBackInvokedCallback,attr"`
		AllowCrossUidActivitySwitchFromBelow string `xml:"allowCrossUidActivitySwitchFromBelow,attr"`
		MetaData                             []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Value    string `xml:"value,attr"`
			Resource string `xml:"resource,attr"`
		} `xml:"meta-data"`
		Profileable struct {
			Text  string `xml:",chardata"`
			Shell string `xml:"shell,attr"`
		} `xml:"profileable"`
		Activity []struct {
			Text                     string `xml:",chardata"`
			Theme                    string `xml:"theme,attr"`
			Label                    string `xml:"label,attr"`
			Name                     string `xml:"name,attr"`
			Exported                 string `xml:"exported,attr"`
			ScreenOrientation        string `xml:"screenOrientation,attr"`
			ConfigChanges            string `xml:"configChanges,attr"`
			WindowSoftInputMode      string `xml:"windowSoftInputMode,attr"`
			TaskAffinity             string `xml:"taskAffinity,attr"`
			ExcludeFromRecents       string `xml:"excludeFromRecents,attr"`
			SupportsPictureInPicture string `xml:"supportsPictureInPicture,attr"`
			LaunchMode               string `xml:"launchMode,attr"`
			NoHistory                string `xml:"noHistory,attr"`
			ShowOnLockScreen         string `xml:"showOnLockScreen,attr"`
			ShowWhenLocked           string `xml:"showWhenLocked,attr"`
			TurnScreenOn             string `xml:"turnScreenOn,attr"`
			RelinquishTaskIdentity   string `xml:"relinquishTaskIdentity,attr"`
			ParentActivityName       string `xml:"parentActivityName,attr"`
			Process                  string `xml:"process,attr"`
			AllowTaskReparenting     string `xml:"allowTaskReparenting,attr"`
			Enabled                  string `xml:"enabled,attr"`
			FitsSystemWindows        string `xml:"fitsSystemWindows,attr"`
			Permission               string `xml:"permission,attr"`
			Icon                     string `xml:"icon,attr"`
			StateNotNeeded           string `xml:"stateNotNeeded,attr"`
			IntentFilter             []struct {
				Text       string `xml:",chardata"`
				AutoVerify string `xml:"autoVerify,attr"`
				Label      string `xml:"label,attr"`
				Action     struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"action"`
				Category []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"category"`
				Data []struct {
					Text        string `xml:",chardata"`
					Scheme      string `xml:"scheme,attr"`
					Host        string `xml:"host,attr"`
					PathPrefix  string `xml:"pathPrefix,attr"`
					Path        string `xml:"path,attr"`
					MimeType    string `xml:"mimeType,attr"`
					PathPattern string `xml:"pathPattern,attr"`
				} `xml:"data"`
			} `xml:"intent-filter"`
			MetaData struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value,attr"`
			} `xml:"meta-data"`
		} `xml:"activity"`
		Receiver []struct {
			Text            string `xml:",chardata"`
			Name            string `xml:"name,attr"`
			Exported        string `xml:"exported,attr"`
			Label           string `xml:"label,attr"`
			InstallLocation string `xml:"installLocation,attr"`
			Permission      string `xml:"permission,attr"`
			Enabled         string `xml:"enabled,attr"`
			DirectBootAware string `xml:"directBootAware,attr"`
			IntentFilter    []struct {
				Text   string `xml:",chardata"`
				Action []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"action"`
				Category struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"category"`
				Data struct {
					Text       string `xml:",chardata"`
					Scheme     string `xml:"scheme,attr"`
					Host       string `xml:"host,attr"`
					PathPrefix string `xml:"pathPrefix,attr"`
				} `xml:"data"`
			} `xml:"intent-filter"`
			MetaData struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name,attr"`
				Resource string `xml:"resource,attr"`
			} `xml:"meta-data"`
		} `xml:"receiver"`
		Service []struct {
			Text                  string `xml:",chardata"`
			Name                  string `xml:"name,attr"`
			Exported              string `xml:"exported,attr"`
			Permission            string `xml:"permission,attr"`
			Label                 string `xml:"label,attr"`
			ForegroundServiceType string `xml:"foregroundServiceType,attr"`
			Enabled               string `xml:"enabled,attr"`
			DirectBootAware       string `xml:"directBootAware,attr"`
			VisibleToInstantApps  string `xml:"visibleToInstantApps,attr"`
			IntentFilter          struct {
				Text     string `xml:",chardata"`
				Priority string `xml:"priority,attr"`
				Action   struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"action"`
				Category struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"category"`
				Data struct {
					Text       string `xml:",chardata"`
					Scheme     string `xml:"scheme,attr"`
					Host       string `xml:"host,attr"`
					PathPrefix string `xml:"pathPrefix,attr"`
				} `xml:"data"`
			} `xml:"intent-filter"`
			MetaData []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value,attr"`
			} `xml:"meta-data"`
		} `xml:"service"`
		Provider []struct {
			Text                string `xml:",chardata"`
			Name                string `xml:"name,attr"`
			Exported            string `xml:"exported,attr"`
			Authorities         string `xml:"authorities,attr"`
			GrantUriPermissions string `xml:"grantUriPermissions,attr"`
			ScreenOrientation   string `xml:"screenOrientation,attr"`
			InitOrder           string `xml:"initOrder,attr"`
			DirectBootAware     string `xml:"directBootAware,attr"`
			Enabled             string `xml:"enabled,attr"`
			MetaData            []struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name,attr"`
				Resource string `xml:"resource,attr"`
				Value    string `xml:"value,attr"`
			} `xml:"meta-data"`
		} `xml:"provider"`
		UsesLibrary []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Required string `xml:"required,attr"`
		} `xml:"uses-library"`
		ActivityAlias struct {
			Text           string `xml:",chardata"`
			Theme          string `xml:"theme,attr"`
			Label          string `xml:"label,attr"`
			Icon           string `xml:"icon,attr"`
			Name           string `xml:"name,attr"`
			Enabled        string `xml:"enabled,attr"`
			Exported       string `xml:"exported,attr"`
			TaskAffinity   string `xml:"taskAffinity,attr"`
			TargetActivity string `xml:"targetActivity,attr"`
			Banner         string `xml:"banner,attr"`
			IntentFilter   struct {
				Text   string `xml:",chardata"`
				Action struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"action"`
				Category []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"category"`
			} `xml:"intent-filter"`
		} `xml:"activity-alias"`
	} `xml:"application"`
}
