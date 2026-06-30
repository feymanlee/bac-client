# BAC Open API Field Audit

Source: https://cloud.baidu.com/doc/ARMCM/s/2kei7tyr3

Generated from the official HTML tables and used to check SDK request structs. Nested fields are included under their parent field where the table exposes them.

| Path | Official request fields |
| --- | --- |
| `/auth/instance/cloud-phone-server-token` | `uuid, instanceCodes, onlineTime, grantControl, CONTROL，可控制；, WATCH：只能观看` |
| `/auth/instance/authorized-server-token` | `uuid, instanceCode, onlineTime, grantControl, CONTROL，可控制；, WATCH：只能观看` |
| `/auth/instance/disconnect` | `uuid, serverToken` |
| `/resources/instance/expire-time-increase` | `serverTokens, time` |
| `/resources/instance/expire-time-update` | `serverTokens, expireTime` |
| `/resources/instance/bind-infos` | `instanceCodes` |
| `/auth/instance/disconnect-all` | `instanceCodes` |
| `/auth/instance/batch-disconnect` | `disconnectList, uuid, serverToken` |
| `/resources/instance/infos` | `page, rows, instanceCodes, merchantPoolNos, deviceIps, deviceIpSegment, instanceType, instanceServerType, romVersion, instanceGradeName, idcCode, usableStatus, maintainStatus, recycleStatus, taskStatus, instanceStatus, deviceStatus, malfunctionStatus, networkStatus, bindStatus, controlStatus, imageVersionId, snapshotMountStatus, snapshotId, tagId, emulationAuthStatus` |
| `/resources/instance/summary-data` | `merchantPoolNo, includeSubPool` |
| `/resources/instance/device-info` | `instanceCodes` |
| `/monitor/pad/available-count.html` | `code, data, totalAvailableCount, idcInfos, idcName, idcCode, availableCount, msg, ts` |
| `/monitor/pad/enable-bind-count.html` | `code, data, totalEnableBindCount, idcInfos, idcName, idcCode, enableBindCount, msg, ts` |
| `/command/pad/screenshot.html` | `padCodes, quality, 其中，1：高质量，2：中质量，3低质量, pictureType` |
| `/command/pad/screenshot-info.html` | `taskIds` |
| `/distribute/task/query-flow` | `instanceCodes, startTime, endTime, billingType, taskDesc` |
| `/distribute/task/query-flow-result` | `taskId` |
| `/resources/instance/file-upload` | `instanceCodes, fileUrl, fileName, fileMd5, autoInstall, 其中，0：不自动安装，1：自动安装, customizeFilePath, 如该参数填/files，将上传至实例的/sdcard/files目录下` |
| `/command/pad/execute-script.html` | `padCodes, scriptContent` |
| `/resources/instance/backup` | `instanceCode, snapshotName, ossConfig, endpoint, bucket, accessKey, secretKey, protocol, snapshotPath, excludes, includes` |
| `/resources/instance/restore` | `snapshotId, instanceCodes` |
| `/resources/instance/snapshot-page` | `snapshotIds, snapshotName, snapshotStatus, wait_create-待创建, creating-创建中, create_success-创建成功, create_failed-创建失败, page, rows` |
| `/auth/instance/authorized-connect` | `uuid, instanceCode, onlineTime, grantControl, CONTROL，可控制；, WATCH：只能观看` |
| `/resources/instance/session-control-switch` | `watchServerTokens, controlServerTokens` |
| `/resources/instance/memory-limit` | `instanceCodes, memoryLimit, isLimit` |
| `/resources/instance/resolution` | `instanceCodes, height, width, fps, dpi` |
| `/resources/instance/install-app` | `appId, instanceCodes` |
| `/resources/instance/uninstall-app` | `appId, instanceCodes` |
| `/resources/app/app-builtin-install` | `appId, instanceCodes` |
| `/resources/app/app-builtin-uninstall` | `appId, instanceCodes` |
| `/command/apps/app-operate.html` | `padCodes, operateType, packageName` |
| `/resources/app/new-version` | `appId, appName, appPackageName, appUrl` |
| `/resources/app/upgrade` | `appId, appName, appPackageName, appUrl` |
| `/distribute/apps/uploads.html` | `apps, appId, appName, pkgName, url, md5sum, taskId` |
| `/resources/app/delete` | `appIds` |
| `/resources/app/page` | `appId, appPackage, page, rows` |
| `/resources/instance/recommend-app-icon-refresh` | `instanceCodes, appIds` |
| `/command/pad/execute-task-info.html` | `taskIds` |
| `/resources/instance/ssh-info` | `instanceCode, connectType, liveTime` |
| `/resources/instance-pool/allocate` | `merchantPoolNo, autoReset, instanceCodes` |
| `/resources/instance/get-screenshot-url` | `instanceCodes, fullQuality, scale, rotate, pictureType` |
| `/resources/instance/destroy-screenshot-url` | `instanceCode` |
| `/resources/instance/set-speed` | `instanceCodes, direction, speed, intranetSpeed` |
| `/resources/device-monitor-info` | `deviceIps` |
| `/resources/device-monitor-info-query` | `deviceIps, startTime, endTime` |
| `/resources/instance-monitor-info` | `instanceCodes` |
| `/resources/instance-app-monitor-info` | `instanceCode, startTime, endTime` |
| `/resources/instance/update-maintain-status` | `instanceCodes, maintainStatus` |
| `/monitor/instance/malfunction-statistics` | `startTime, endTime, timeUnit, includeSubMerchant` |
| `/resources/instance/batch-execute-script` | `scripts, instanceCode, scriptContent` |
| `/resources/instance/file-download` | `instanceFiles, instanceCode, filePath` |
| `/resources/instance/event-task-stop` | `taskIds` |
| `/resources/device/page` | `page, rows, deviceCodes, deviceIps, deviceIpSegment, idcCode, instanceServerType, romVersion, deviceStatus` |
| `/monitor/network-bandwidth/list` | `idcCode, statUnit, beginTime, endTime` |
| `/resources/instance/network-proxy-workflow-create` | `networkProxyConfigs, instanceCode, proxyHost, proxyPort, proxyUser, proxyPassword, proxyWhite` |
| `/resources/instance/data-copy` | `dataCopyList, sourceInstanceCode, targetInstanceCode, imei, serialno, wifimac, androidid, model, brand, manufacturer, includes, 此目录需要对安卓目录结构，非常了解，如果客户未做这方面调研，建议设置为空字符串, 此目录为/data下的相对目录，不支持非/data目录下的include, *符号表示此目录下所有层级子目录，如果只有一个，只会压缩一层目录, 如果未填写includes，或者includes为空，则默认打包所有/data下的目录, excludes, 同上，此目录为/data下的相对目录。可以使用通配符*, 如果未填写excludes，或者excludes为空，则默认不排除任务目录, reset` |
| `/resources/instance-tag/page` | `tagName, tagId, page, rows` |
| `/resources/instance-tag/save` | `tagName, remark` |
| `/resources/instance-tag/update` | `tagId, tagName, remark` |
| `/resources/instance-tag/remove` | `tagIds` |
| `/resources/instance-tag/relate-add` | `tagIds, instanceCodes` |
| `/resources/instance-tag/relate-remove` | `tagIds, instanceCodes` |
| `/resources/instance/custom-code-update` | `instanceCustomList, instanceCode, customCode` |
| `/merchant/alarm-strategy/page` | `page, rows, alarmStrategyName, alarmResourceType, 1. OSS 对象存储、, 2. DUFS DuFS存储、, 3. BANDWIDTH 带宽用量, enableStatus` |
| `/merchant/alarm-strategy/add` | `alarmStrategyName, alarmStrategyDesc, alarmResourceType, 1. OSS 对象存储, 2. DUFS DuFS存储, 3. BANDWIDTH 带宽用量, idcCodes, 一个机房可以创建1个带宽用量告警，创建5个对象存储、DuFS存储, alarmMetrics, 1. remaining-percentage 剩余百分比 单位 %, 2. used-percentage 已用量百分比 单位 %, 3. bandwidth-peak 当月带宽峰值 单位Mbps, alarmThreshold, 当月带宽峰值范围[10, 99999], 剩余百分比和已用量百分比范围[1, 99], alarmSilencePeriod, smsNotifyStatus, smsNotifyObjects, 同一手机号码每分钟最多能接收的短信通知3个，每小时20个，每天能接收的短信通知最多100个, callbackNotifyStatus` |
| `/merchant/alarm-strategy/update` | `alarmStrategyId, alarmStrategyName, alarmStrategyDesc, alarmThreshold, 当月带宽峰值范围[10, 99999], 剩余百分比和已用量百分比范围[1, 99], alarmSilencePeriod, smsNotifyStatus, smsNotifyObjects, 同一手机号码每分钟最多能接收的短信通知3个，每小时20个，每天能接收的短信通知最多100个, callbackNotifyStatus` |
| `/merchant/alarm-strategy/remove` | `alarmStrategyIds` |
| `/merchant/alarm-strategy/update-enable-status` | `alarmStrategyIds, enableStatus` |
| `/resources/app/get-desktop-icon-config` | `instanceCodes, appIds, container, screen, x, y` |
| `/resources/app/save-desktop-icon-config` | `instanceCodes, appId, container, screen, x, y, overwriteCoordinate` |
| `/resources/app/remove-desktop-icon-config` | `instanceCodes, appIds` |
| `/resources/instance/export-ip` | `instanceCodes` |
| `/resources/instance/app-install-list` | `instanceCodes` |
| `/resources/instance-pool/save` | `parentMerchantPoolNo, instancePoolName, instancePoolType` |
| `/merchant/open-account/save` | `userName, phone, nickname, roleNames, merchantPoolNos` |
| `/merchant/sub-merchant/add` | `parentMerchantCode, merchantCode, merchantName, merchantType, merchantPhone, adminUserName, adminPhone` |
| `/resources/instance/deploy-marketing-suite` | `instanceCodes, authCodes, 如果实例已绑定该应用，将不消耗授权码，该参数允许为空, appPackage, useMerchantAuthCode, 当传入true时，对于未绑定指定应用的实例，将自动使用商户下的空闲的授权码进行绑定，无需传入authCodes参数` |
| `/resources/instance/metric-detail` | `instanceCodes, recordTime` |
| `/command/pad/execute-task-type-list` | `code, data, type, name, msg, ts` |
| `/command/pad/execute-task-page` | `taskIds, taskType, page, rows` |
| `/resources/device/clean-app-cache` | `DeviceIps` |
| `/resources/emulation/auth-code-page` | `authCodes, authStatus, instanceCodes, packageName, page, rows` |
| `/resources/instance-base-image/list` | `instanceServerType, romVersion, android6.0, android8.1, android10.0, android12.0, android13.0, imageVersionIds, page, rows` |
| `/resources/instance-image/list` | `page, rows, imageVersionId, imageVersionIds, instanceServerType, romVersion, android6.0, android8.1, android10.0, android12.0, android13.0, imageVersionName` |
| `/resources/instance-image/upload` | `imageFiles, imageFileUrl, imageFileName, 镜像⽂件名称需要以镜像⽂件类型为前缀，⽐如说root_aosp类型的就需要以root_aosp开头命名，如root_aosp_xx.img ；, 如果是安卓12，在⽂件类型后⾯应该带f2fs，如system_aosp_f2fs_xxx.img, imageFileType, root_aosp, system_aosp, vendor_aosp, super_aosp, imageFileMd5, instanceServerType, romVersion, android6.0, android8.1, android10.0, android12.0, android13.0, baseImageVersionId, imageVersionName, describe` |
| `/resources/instance-image/upload-info` | `imageVersionId` |
| `/resources/instance-image/remove` | `imageVersionIds` |
| `/resources/instance-image/update` | `imageVersionId, instanceCodes, configId, resourcePackageId, reset, autoInstall` |
| `/resources/instance-image/update-info` | `taskIds` |
| `/resources/device/image-update` | `imageVersionId, deviceIps, configId, resourcePackageId, reset, autoInstall` |
| `/command/pad/reboot.html` | `padCodes, merchantPoolNos` |
| `/distribute/pad/new-pad.html` | `padModels, padCode, imei, serialno, wifimac, androidid, model, brand, manufacturer` |
| `/command/pad/reset.html` | `padCodes, merchantPoolNos, resetType, IMAGE_ONLY：支持只恢复镜像, DATA_AND_IMAGE：支持同时恢复镜像、清理数据` |
| `/resources/instance/reboot-remote-play` | `instanceCodes` |
| `/resources/device/reboot` | `deviceCodes, rebootType` |
| `/resources/dufs-snapshot/init` | `instanceCodes, quotaCapacity, memoryLimit` |
| `/resources/dufs-snapshot/unmount` | `instanceCodes, snapshotName` |
| `/resources/dufs-snapshot/mount` | `instanceCodes, snapshotId, quotaCapacity, memoryLimit` |
| `/resources/dufs-snapshot/page` | `snapshotName, snapshotStatus, snapshotIds, page, rows` |
| `/resources/dufs-snapshot/remove` | `snapshotIds` |
| `/resources/dufs-snapshot/batch-mount` | `snapshotMountInfos, instanceCode, snapshotId, quotaCapacity, memoryLimit` |
| `/resources/dufs-snapshot/quota-set` | `instanceCodes, quotaCapacity` |
