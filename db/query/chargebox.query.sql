-- name: CreateChargeBox :one
BEGIN;

INSERT INTO `chargebox` (
    `chargeBoxId`,
    `endpoint_address`,
    `chargePointVendor`,
    `chargePointModel`
) VALUES ($1, $2, $3, $4);

-- Retrieve the inserted values
SELECT
    `chargeBoxId`,
    `endpoint_address`,
    `chargePointVendor`,
    `chargePointModel`
FROM
    `chargebox`
WHERE
    `chargeBoxId` = LAST_INSERT_ID();

COMMIT;
