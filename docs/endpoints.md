# BAC Open API Endpoint Coverage

Source: https://cloud.baidu.com/doc/ARMCM/s/2kei7tyr3

This page lists the BAC Open API paths covered by typed SDK methods. All endpoints use `POST` and the encrypted request envelope handled by `Client.Do`.

For request/response modeling rules, see [API Conventions](api-conventions.md).

## Coverage

| Group | Path | Method | Response shape |
| --- | --- | --- | --- |
| `Auth` | `/auth/instance/cloud-phone-server-token` | `GetServerToken` | `ServerTokenResponse` |
| `Auth` | `/auth/instance/authorized-server-token` | `GetAuthorizedServerToken` | `ServerTokenResponse` |
| `Auth` | `/auth/instance/authorized-connect` | `AuthorizedConnect` | `AuthorizedConnectResponse` |
| `Auth` | `/auth/instance/disconnect` | `DisconnectInstance` | `empty` |
| `Auth` | `/auth/instance/batch-disconnect` | `BatchDisconnectInstances` | `BatchDisconnectResponse` |
| `Auth` | `/auth/instance/disconnect-all` | `DisconnectAllInstances` | `empty` |
| `Instance` | `/resources/instance/infos` | `ListInstances` | `Page[Instance]` |
| `Instance` | `/resources/instance/summary-data` | `GetSummaryData` | `RawObject` |
| `Instance` | `/resources/instance/device-info` | `GetInstanceDeviceInfo` | `[]DeviceInfo` |
| `Instance` | `/resources/instance/file-upload` | `UploadFile` | `[]TaskResponse` |
| `Instance` | `/resources/instance/file-download` | `DownloadFile` | `BatchTaskResponse` |
| `Instance` | `/resources/instance/backup` | `BackupInstance` | `BackupInstanceResponse` |
| `Instance` | `/resources/instance/restore` | `RestoreInstance` | `[]TaskResponse` |
| `Instance` | `/resources/instance/snapshot-page` | `ListInstanceSnapshots` | `Page[SnapshotInfo]` |
| `Instance` | `/resources/instance/session-control-switch` | `SwitchSessionControl` | `empty` |
| `Instance` | `/resources/instance/memory-limit` | `SetMemoryLimit` | `BatchTaskResponse` |
| `Instance` | `/resources/instance/resolution` | `SetResolution` | `[]TaskResponse` |
| `Instance` | `/resources/instance/install-app` | `InstallApp` | `[]TaskResponse` |
| `Instance` | `/resources/instance/uninstall-app` | `UninstallApp` | `[]TaskResponse` |
| `Instance` | `/resources/instance/reboot-remote-play` | `RebootInstance` | `TaskResponse` |
| `Instance` | `/resources/instance/batch-execute-script` | `BatchExecuteScript` | `BatchTaskResponse` |
| `Instance` | `/resources/instance/bind-infos` | `GetBindInfos` | `[]BindInfo` |
| `Instance` | `/resources/instance/custom-code-update` | `UpdateCustomCode` | `BatchTaskResponse` |
| `Instance` | `/resources/instance/data-copy` | `CopyInstanceData` | `[]TaskResponse` |
| `Instance` | `/resources/instance/deploy-marketing-suite` | `DeployMarketingSuite` | `[]TaskResponse` |
| `Instance` | `/resources/instance/destroy-screenshot-url` | `DestroyScreenshotURL` | `empty` |
| `Instance` | `/resources/instance/event-task-stop` | `StopEventTasks` | `empty` |
| `Instance` | `/resources/instance/expire-time-increase` | `IncreaseExpireTime` | `empty` |
| `Instance` | `/resources/instance/expire-time-update` | `UpdateExpireTime` | `empty` |
| `Instance` | `/resources/instance/export-ip` | `ExportIP` | `RawObject` |
| `Instance` | `/resources/instance/get-screenshot-url` | `GetScreenshotURL` | `ScreenshotURLResponse` |
| `Instance` | `/resources/instance/network-proxy-workflow-create` | `CreateNetworkProxyWorkflow` | `[]TaskResponse` |
| `Instance` | `/resources/instance/set-speed` | `SetSpeed` | `BatchTaskResponse` |
| `Instance` | `/resources/instance/ssh-info` | `GetSSHInfo` | `SSHInfoResponse` |
| `Instance` | `/resources/instance/update-maintain-status` | `UpdateMaintainStatus` | `empty` |
| `Command` | `/command/pad/screenshot.html` | `Screenshot` | `[]TaskResponse` |
| `Command` | `/command/pad/screenshot-info.html` | `GetScreenshotResult` | `[]TaskResult` |
| `Command` | `/command/apps/app-operate.html` | `OperateApp / StartApp / StopApp` | `[]TaskResponse` |
| `Command` | `/command/pad/execute-script.html` | `ExecuteScript` | `[]TaskResponse` |
| `Command` | `/command/pad/execute-task-info.html` | `GetTaskResult` | `[]TaskResult` |
| `Command` | `/command/pad/execute-task-page` | `ListTaskResults` | `Page[TaskResult]` |
| `Command` | `/command/pad/execute-task-type-list` | `ListTaskTypes` | `[]TaskType` |
| `Command` | `/command/pad/reboot.html` | `CommandReboot` | `TaskResponse` |
| `Command` | `/command/pad/reset.html` | `CommandReset / ResetDevice` | `TaskResponse` |
| `Command` | `/distribute/pad/new-pad.html` | `NewPad` | `[]TaskResponse` |
| `Device` | `/resources/device/page` | `ListDevices` | `Page[RawObject]` |
| `Device` | `/resources/device/reboot` | `RebootDevice` | `[]TaskResponse` |
| `Device` | `/resources/device/clean-app-cache` | `CleanAppCache` | `[]TaskResponse` |
| `Device` | `/resources/device/image-update` | `UpdateDeviceImage` | `[]DeviceImageUpdateResult` |
| `App` | `/distribute/apps/uploads.html` | `UploadApps` | `empty` |
| `App` | `/resources/app/delete` | `DeleteApp` | `empty` |
| `App` | `/resources/app/page` | `ListApps` | `Page[AppInfo]` |
| `App` | `/resources/app/new-version` | `NewAppVersion` | `empty` |
| `App` | `/resources/app/upgrade` | `UpgradeApp` | `empty` |
| `App` | `/resources/app/app-builtin-install` | `BuiltinInstallApp` | `BatchTaskResponse` |
| `App` | `/resources/app/app-builtin-uninstall` | `BuiltinUninstallApp` | `BatchTaskResponse` |
| `App` | `/resources/app/get-desktop-icon-config` | `GetDesktopIconConfig` | `[]DesktopIconConfig` |
| `App` | `/resources/app/save-desktop-icon-config` | `SaveDesktopIconConfig` | `empty` |
| `App` | `/resources/app/remove-desktop-icon-config` | `RemoveDesktopIconConfig` | `empty` |
| `App` | `/resources/instance/recommend-app-icon-refresh` | `RefreshRecommendAppIcons` | `BatchTaskResponse` |
| `App` | `/resources/instance/app-install-list` | `ListInstalledApps` | `[]InstalledAppList` |
| `Distribute` | `/distribute/task/query-flow` | `QueryFlow` | `RawObject` |
| `Distribute` | `/distribute/task/query-flow-result` | `QueryFlowResult` | `RawObject` |
| `Monitor` | `/monitor/instance/malfunction-statistics` | `GetMalfunctionStatistics` | `[]MalfunctionStatistic` |
| `Monitor` | `/monitor/network-bandwidth/list` | `ListNetworkBandwidth` | `Page[RawObject]` |
| `Monitor` | `/monitor/pad/available-count.html` | `GetAvailablePadCount` | `RawObject` |
| `Monitor` | `/monitor/pad/enable-bind-count.html` | `GetEnableBindPadCount` | `RawObject` |
| `Monitor` | `/resources/device-monitor-info` | `GetDeviceMonitorInfo` | `[]DeviceMonitorInfo` |
| `Monitor` | `/resources/device-monitor-info-query` | `QueryDeviceMonitorInfo` | `[]RawObject` |
| `Monitor` | `/resources/instance-monitor-info` | `GetInstanceMonitorInfo` | `[]RawObject` |
| `Monitor` | `/resources/instance-app-monitor-info` | `GetInstanceAppMonitorInfo` | `[]RawObject` |
| `Monitor` | `/resources/instance/metric-detail` | `GetInstanceMetricDetail` | `RawObject` |
| `Merchant` | `/merchant/alarm-strategy/add` | `AddAlarmStrategy` | `empty` |
| `Merchant` | `/merchant/alarm-strategy/page` | `ListAlarmStrategies` | `Page[AlarmStrategy]` |
| `Merchant` | `/merchant/alarm-strategy/remove` | `RemoveAlarmStrategy` | `empty` |
| `Merchant` | `/merchant/alarm-strategy/update` | `UpdateAlarmStrategy` | `empty` |
| `Merchant` | `/merchant/alarm-strategy/update-enable-status` | `UpdateAlarmStrategyEnableStatus` | `empty` |
| `Merchant` | `/merchant/open-account/save` | `OpenAccount` | `OpenAccountResponse` |
| `Merchant` | `/merchant/sub-merchant/add` | `AddSubMerchant` | `AddSubMerchantResponse` |
| `Pool` | `/resources/instance-pool/allocate` | `AllocateInstancePool` | `empty` |
| `Pool` | `/resources/instance-pool/save` | `SaveInstancePool` | `SaveInstancePoolResponse` |
| `Tag` | `/resources/instance-tag/page` | `ListInstanceTags` | `Page[InstanceTag]` |
| `Tag` | `/resources/instance-tag/save` | `SaveInstanceTag` | `empty` |
| `Tag` | `/resources/instance-tag/update` | `UpdateInstanceTag` | `empty` |
| `Tag` | `/resources/instance-tag/remove` | `RemoveInstanceTag` | `empty` |
| `Tag` | `/resources/instance-tag/relate-add` | `AddInstanceTagRelation` | `empty` |
| `Tag` | `/resources/instance-tag/relate-remove` | `RemoveInstanceTagRelation` | `empty` |
| `Image` | `/resources/instance-base-image/list` | `ListBaseImages` | `Page[ImageInfo]` |
| `Image` | `/resources/instance-image/list` | `ListInstanceImages` | `Page[ImageInfo]` |
| `Image` | `/resources/instance-image/upload` | `UploadImage` | `UploadImageResponse` |
| `Image` | `/resources/instance-image/upload-info` | `GetImageUploadInfo` | `ImageUploadInfo` |
| `Image` | `/resources/instance-image/remove` | `RemoveImage` | `empty` |
| `Image` | `/resources/instance-image/update` | `UpdateImage` | `[]TaskResponse` |
| `Image` | `/resources/instance-image/update-info` | `GetImageUpdateInfo` | `[]ImageUpdateInfo` |
| `Snapshot` | `/resources/dufs-snapshot/init` | `InitDUFSSnapshot` | `[]TaskResponse` |
| `Snapshot` | `/resources/dufs-snapshot/unmount` | `UnmountDUFSSnapshot` | `[]TaskResponse` |
| `Snapshot` | `/resources/dufs-snapshot/mount` | `MountDUFSSnapshot` | `[]TaskResponse` |
| `Snapshot` | `/resources/dufs-snapshot/batch-mount` | `BatchMountDUFSSnapshot` | `[]TaskResponse` |
| `Snapshot` | `/resources/dufs-snapshot/page` | `ListDUFSSnapshots` | `Page[SnapshotInfo]` |
| `Snapshot` | `/resources/dufs-snapshot/remove` | `RemoveDUFSSnapshot` | `empty` |
| `Snapshot` | `/resources/dufs-snapshot/quota-set` | `SetDUFSSnapshotQuota` | `[]TaskResponse` |
| `Emulation` | `/resources/emulation/auth-code-page` | `ListEmulationAuthCodes` | `Page[RawObject]` |
