CREATE TABLE `chargebox` (
  `chargeBoxId` varchar(30) NOT NULL,
  `endpoint_address` varchar(45) DEFAULT NULL,
  `ocppProtocol` varchar(10) DEFAULT NULL,
  `chargePointVendor` varchar(20) DEFAULT NULL,
  `chargePointModel` varchar(20) DEFAULT NULL,
  `chargePointSerialNumber` varchar(25) DEFAULT NULL,
  `chargeBoxSerialNumber` varchar(25) DEFAULT NULL,
  `fwVersion` varchar(50) DEFAULT NULL,
  `fwUpdateStatus` varchar(25) DEFAULT NULL,
  `fwUpdateTimestamp` timestamp NULL DEFAULT NULL,
  `iccid` varchar(20) DEFAULT NULL,
  `imsi` varchar(20) DEFAULT NULL,
  `meterType` varchar(25) DEFAULT NULL,
  `meterSerialNumber` varchar(25) DEFAULT NULL,
  `diagnosticsStatus` varchar(20) DEFAULT NULL,
  `diagnosticsTimestamp` timestamp NULL DEFAULT NULL,
  `lastHeartbeatTimestamp` timestamp NULL DEFAULT NULL,
  `note` text DEFAULT NULL,
  PRIMARY KEY (`chargeBoxId`),
  UNIQUE KEY `chargeBoxId_UNIQUE` (`chargeBoxId`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci