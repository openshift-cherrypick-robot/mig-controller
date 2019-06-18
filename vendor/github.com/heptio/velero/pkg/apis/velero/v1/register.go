/*
Copyright 2017 the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeBuilder collects the scheme builder functions for the Velero API
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// AddToScheme applies the SchemeBuilder functions to a specified scheme
	AddToScheme = SchemeBuilder.AddToScheme
)

// GroupName is the group name for the Velero API
const GroupName = "velero.io"

// SchemeGroupVersion is the GroupVersion for the Velero API
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

// Resource gets a Velero GroupResource for a specified resource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

type typeInfo struct {
	PluralName   string
	ItemType     runtime.Object
	ItemListType runtime.Object
}

func newTypeInfo(pluralName string, itemType, itemListType runtime.Object) typeInfo {
	return typeInfo{
		PluralName:   pluralName,
		ItemType:     itemType,
		ItemListType: itemListType,
	}
}

// CustomResources returns a map of all custom resources within the Velero
// API group, keyed on Kind.
func CustomResources() map[string]typeInfo {
	return map[string]typeInfo{
		"Backup":                 newTypeInfo("backups", &Backup{}, &BackupList{}),
		"Restore":                newTypeInfo("restores", &Restore{}, &RestoreList{}),
		"Schedule":               newTypeInfo("schedules", &Schedule{}, &ScheduleList{}),
		"DownloadRequest":        newTypeInfo("downloadrequests", &DownloadRequest{}, &DownloadRequestList{}),
		"DeleteBackupRequest":    newTypeInfo("deletebackuprequests", &DeleteBackupRequest{}, &DeleteBackupRequestList{}),
		"PodVolumeBackup":        newTypeInfo("podvolumebackups", &PodVolumeBackup{}, &PodVolumeBackupList{}),
		"PodVolumeRestore":       newTypeInfo("podvolumerestores", &PodVolumeRestore{}, &PodVolumeRestoreList{}),
		"ResticRepository":       newTypeInfo("resticrepositories", &ResticRepository{}, &ResticRepositoryList{}),
		"BackupStorageLocation":  newTypeInfo("backupstoragelocations", &BackupStorageLocation{}, &BackupStorageLocationList{}),
		"VolumeSnapshotLocation": newTypeInfo("volumesnapshotlocations", &VolumeSnapshotLocation{}, &VolumeSnapshotLocationList{}),
		"ServerStatusRequest":    newTypeInfo("serverstatusrequests", &ServerStatusRequest{}, &ServerStatusRequestList{}),
	}
}

func addKnownTypes(scheme *runtime.Scheme) error {
	for _, typeInfo := range CustomResources() {
		scheme.AddKnownTypes(SchemeGroupVersion, typeInfo.ItemType, typeInfo.ItemListType)
	}

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
