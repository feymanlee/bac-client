# BAC Open API Endpoint Coverage

Source: https://cloud.baidu.com/doc/ARMCM/s/2kei7tyr3

All endpoints use `POST` and the encrypted envelope handled by `Client.Do`.

## Typed Wrappers Implemented

| Group | Path | Method |
| --- | --- | --- |
| Auth | `/auth/instance/cloud-phone-server-token` | `GetServerToken` |
| Auth | `/auth/instance/authorized-server-token` | `GetAuthorizedServerToken` |
| Auth | `/auth/instance/authorized-connect` | `AuthorizedConnect` |
| Auth | `/auth/instance/disconnect` | `DisconnectInstance` |
| Auth | `/auth/instance/batch-disconnect` | `BatchDisconnectInstances` |
| Auth | `/auth/instance/disconnect-all` | `DisconnectAllInstances` |
| Instance | `/resources/instance/infos` | `ListInstances` |
| Instance | `/resources/instance/device-info` | `GetInstanceDeviceInfo` |
| Instance | `/resources/instance/file-upload` | `UploadFile` |
| Instance | `/resources/instance/backup` | `BackupInstance` |
| Instance | `/resources/instance/restore` | `RestoreInstance` |
| Instance | `/resources/instance/session-control-switch` | `SwitchSessionControl` |
| Instance | `/resources/instance/resolution` | `SetResolution` |
| Instance | `/resources/instance/install-app` | `InstallApp` |
| Instance | `/resources/instance/uninstall-app` | `UninstallApp` |
| Instance | `/resources/instance/reboot-remote-play` | `RebootInstance` |
| Instance | `/resources/device/reboot` | `RebootDevice` |
| Instance | `/distribute/pad/new-pad.html` | `NewPad` |
| Instance | `/resources/device/clean-app-cache` | `CleanAppCache` |
| Command | `/command/pad/screenshot.html` | `Screenshot` |
| Command | `/command/pad/screenshot-info.html` | `GetScreenshotResult` |
| Command | `/command/apps/app-operate.html` | `OperateApp`, `StartApp`, `StopApp` |
| Command | `/command/pad/execute-task-info.html` | `GetTaskResult` |
| Command | `/command/pad/execute-task-page` | `ListTaskResults` |
| Command | `/command/pad/execute-task-type-list` | `ListTaskTypes` |
| Command | `/command/pad/reboot.html` | `CommandReboot` |
| Command | `/command/pad/reset.html` | `CommandReset`, `ResetDevice` |
| Image | `/resources/instance-base-image/list` | `ListBaseImages` |
| Image | `/resources/instance-image/list` | `ListInstanceImages` |
| Image | `/resources/instance-image/upload` | `UploadImage` |
| Image | `/resources/instance-image/upload-info` | `GetImageUploadInfo` |
| Image | `/resources/instance-image/update` | `UpdateImage` |
| Image | `/resources/instance-image/update-info` | `GetImageUpdateInfo` |
| Image | `/resources/instance-image/remove` | `RemoveImage` |

## Additional Typed Wrappers Implemented

These endpoints now have typed request structs and method wrappers. Some responses intentionally expose `RawObject` or `Page[RawObject]` where the official response shape is broad or likely to evolve.

| Group | Path | Method |
| --- | --- | --- |
| Command | `/command/pad/execute-script.html` | `ExecuteScript` |
| Distribute | `/distribute/apps/uploads.html` | `UploadApps` |
| Distribute | `/distribute/task/query-flow` | `QueryFlow` |
| Distribute | `/distribute/task/query-flow-result` | `QueryFlowResult` |
| Merchant | `/merchant/alarm-strategy/add` | `AddAlarmStrategy` |
| Merchant | `/merchant/alarm-strategy/page` | `ListAlarmStrategies` |
| Merchant | `/merchant/alarm-strategy/remove` | `RemoveAlarmStrategy` |
| Merchant | `/merchant/alarm-strategy/update` | `UpdateAlarmStrategy` |
| Merchant | `/merchant/alarm-strategy/update-enable-status` | `UpdateAlarmStrategyEnableStatus` |
| Merchant | `/merchant/open-account/save` | `OpenAccount` |
| Merchant | `/merchant/sub-merchant/add` | `AddSubMerchant` |
| Monitor | `/monitor/instance/malfunction-statistics` | `GetMalfunctionStatistics` |
| Monitor | `/monitor/network-bandwidth/list` | `ListNetworkBandwidth` |
| Monitor | `/monitor/pad/available-count.html` | `GetAvailablePadCount` |
| Monitor | `/monitor/pad/enable-bind-count.html` | `GetEnableBindPadCount` |
| App | `/resources/app/app-builtin-install` | `BuiltinInstallApp` |
| App | `/resources/app/app-builtin-uninstall` | `BuiltinUninstallApp` |
| App | `/resources/app/delete` | `DeleteApp` |
| App | `/resources/app/get-desktop-icon-config` | `GetDesktopIconConfig` |
| App | `/resources/app/new-version` | `NewAppVersion` |
| App | `/resources/app/page` | `ListApps` |
| App | `/resources/app/remove-desktop-icon-config` | `RemoveDesktopIconConfig` |
| App | `/resources/app/save-desktop-icon-config` | `SaveDesktopIconConfig` |
| App | `/resources/app/upgrade` | `UpgradeApp` |
| Monitor | `/resources/device-monitor-info` | `GetDeviceMonitorInfo` |
| Monitor | `/resources/device-monitor-info-query` | `QueryDeviceMonitorInfo` |
| Device | `/resources/device/image-update` | `UpdateDeviceImage` (`DeviceImageUpdateRequest` -> `[]DeviceImageUpdateResult`) |
| Snapshot | `/resources/dufs-snapshot/batch-mount` | `BatchMountDUFSSnapshot` |
| Snapshot | `/resources/dufs-snapshot/init` | `InitDUFSSnapshot` |
| Snapshot | `/resources/dufs-snapshot/mount` | `MountDUFSSnapshot` |
| Snapshot | `/resources/dufs-snapshot/page` | `ListDUFSSnapshots` |
| Snapshot | `/resources/dufs-snapshot/quota-set` | `SetDUFSSnapshotQuota` |
| Snapshot | `/resources/dufs-snapshot/remove` | `RemoveDUFSSnapshot` |
| Snapshot | `/resources/dufs-snapshot/unmount` | `UnmountDUFSSnapshot` |
| Monitor | `/resources/instance-app-monitor-info` | `GetInstanceAppMonitorInfo` |
| Monitor | `/resources/instance-monitor-info` | `GetInstanceMonitorInfo` |
| Pool | `/resources/instance-pool/allocate` | `AllocateInstancePool` |
| Pool | `/resources/instance-pool/save` | `SaveInstancePool` |
| Tag | `/resources/instance-tag/page` | `ListInstanceTags` |
| Tag | `/resources/instance-tag/relate-add` | `AddInstanceTagRelation` |
| Tag | `/resources/instance-tag/relate-remove` | `RemoveInstanceTagRelation` |
| Tag | `/resources/instance-tag/remove` | `RemoveInstanceTag` |
| Tag | `/resources/instance-tag/save` | `SaveInstanceTag` |
| Tag | `/resources/instance-tag/update` | `UpdateInstanceTag` |
| App | `/resources/instance/app-install-list` | `ListInstalledApps` |
| Instance | `/resources/instance/batch-execute-script` | `BatchExecuteScript` |
| Instance | `/resources/instance/bind-infos` | `GetBindInfos` |
| Instance | `/resources/instance/custom-code-update` | `UpdateCustomCode` |
| Instance | `/resources/instance/data-copy` | `CopyInstanceData` |
| Instance | `/resources/instance/deploy-marketing-suite` | `DeployMarketingSuite` |
| Instance | `/resources/instance/destroy-screenshot-url` | `DestroyScreenshotURL` |
| Instance | `/resources/instance/event-task-stop` | `StopEventTasks` |
| Instance | `/resources/instance/expire-time-increase` | `IncreaseExpireTime` |
| Instance | `/resources/instance/expire-time-update` | `UpdateExpireTime` |
| Instance | `/resources/instance/export-ip` | `ExportIP` |
| Instance | `/resources/instance/file-download` | `DownloadFile` |
| Instance | `/resources/instance/get-screenshot-url` | `GetScreenshotURL` |
| Instance | `/resources/instance/memory-limit` | `SetMemoryLimit` |
| Monitor | `/resources/instance/metric-detail` | `GetInstanceMetricDetail` |
| Instance | `/resources/instance/network-proxy-workflow-create` | `CreateNetworkProxyWorkflow` |
| App | `/resources/instance/recommend-app-icon-refresh` | `RefreshRecommendAppIcons` |
| Instance | `/resources/instance/set-speed` | `SetSpeed` |
| Snapshot | `/resources/instance/snapshot-page` | `ListInstanceSnapshots` |
| Instance | `/resources/instance/ssh-info` | `GetSSHInfo` |
| Instance | `/resources/instance/summary-data` | `GetSummaryData` |
| Instance | `/resources/instance/update-maintain-status` | `UpdateMaintainStatus` |

## Remaining Low-Level Only

| Path | Status |
| --- | --- |
| `/resources/emulation/auth-code-page` | `ListEmulationAuthCodes` |
