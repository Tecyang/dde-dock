// This file is automatically generated, please don't edit manully.
package main

// Get key type
func getSettingConnectionKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_CONNECTION_ID:
		t = ktypeString
	case NM_SETTING_CONNECTION_UUID:
		t = ktypeString
	case NM_SETTING_CONNECTION_TYPE:
		t = ktypeString
	case NM_SETTING_CONNECTION_PERMISSIONS:
		t = ktypeArrayString
	case NM_SETTING_CONNECTION_AUTOCONNECT:
		t = ktypeBoolean
	case NM_SETTING_CONNECTION_TIMESTAMP:
		t = ktypeUint64
	case NM_SETTING_CONNECTION_READ_ONLY:
		t = ktypeBoolean
	case NM_SETTING_CONNECTION_ZONE:
		t = ktypeString
	case NM_SETTING_CONNECTION_MASTER:
		t = ktypeString
	case NM_SETTING_CONNECTION_SLAVE_TYPE:
		t = ktypeString
	case NM_SETTING_CONNECTION_SECONDARIES:
		t = ktypeArrayString
	}
	return
}

// Get key's default value
func getSettingConnectionKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		LOGGER.Error("invalid key:", key)
	case NM_SETTING_CONNECTION_ID:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_UUID:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_TYPE:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_PERMISSIONS:
		valueJSON = `null`
	case NM_SETTING_CONNECTION_AUTOCONNECT:
		valueJSON = `true`
	case NM_SETTING_CONNECTION_TIMESTAMP:
		valueJSON = `0`
	case NM_SETTING_CONNECTION_READ_ONLY:
		valueJSON = `false`
	case NM_SETTING_CONNECTION_ZONE:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_MASTER:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_SLAVE_TYPE:
		valueJSON = `""`
	case NM_SETTING_CONNECTION_SECONDARIES:
		valueJSON = `null`
	}
	return
}

// Get JSON value generally
func generalGetSettingConnectionKeyJSON(data _ConnectionData, key string) (value string) {
	switch key {
	default:
		LOGGER.Error("generalGetSettingConnectionKeyJSON: invalide key", key)
	case NM_SETTING_CONNECTION_ID:
		value = getSettingConnectionIdJSON(data)
	case NM_SETTING_CONNECTION_UUID:
		value = getSettingConnectionUuidJSON(data)
	case NM_SETTING_CONNECTION_TYPE:
		value = getSettingConnectionTypeJSON(data)
	case NM_SETTING_CONNECTION_PERMISSIONS:
		value = getSettingConnectionPermissionsJSON(data)
	case NM_SETTING_CONNECTION_AUTOCONNECT:
		value = getSettingConnectionAutoconnectJSON(data)
	case NM_SETTING_CONNECTION_TIMESTAMP:
		value = getSettingConnectionTimestampJSON(data)
	case NM_SETTING_CONNECTION_READ_ONLY:
		value = getSettingConnectionReadOnlyJSON(data)
	case NM_SETTING_CONNECTION_ZONE:
		value = getSettingConnectionZoneJSON(data)
	case NM_SETTING_CONNECTION_MASTER:
		value = getSettingConnectionMasterJSON(data)
	case NM_SETTING_CONNECTION_SLAVE_TYPE:
		value = getSettingConnectionSlaveTypeJSON(data)
	case NM_SETTING_CONNECTION_SECONDARIES:
		value = getSettingConnectionSecondariesJSON(data)
	}
	return
}

// Getter
func getSettingConnectionId(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID).(string)
	return
}
func getSettingConnectionUuid(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID).(string)
	return
}
func getSettingConnectionType(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE).(string)
	return
}
func getSettingConnectionPermissions(data _ConnectionData) (value []string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS).([]string)
	return
}
func getSettingConnectionAutoconnect(data _ConnectionData) (value bool) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT).(bool)
	return
}
func getSettingConnectionTimestamp(data _ConnectionData) (value uint64) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP).(uint64)
	return
}
func getSettingConnectionReadOnly(data _ConnectionData) (value bool) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY).(bool)
	return
}
func getSettingConnectionZone(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE).(string)
	return
}
func getSettingConnectionMaster(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER).(string)
	return
}
func getSettingConnectionSlaveType(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE).(string)
	return
}
func getSettingConnectionSecondaries(data _ConnectionData) (value []string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES).([]string)
	return
}

// Setter
func setSettingConnectionId(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID, value)
}
func setSettingConnectionUuid(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID, value)
}
func setSettingConnectionType(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE, value)
}
func setSettingConnectionPermissions(data _ConnectionData, value []string) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS, value)
}
func setSettingConnectionAutoconnect(data _ConnectionData, value bool) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT, value)
}
func setSettingConnectionTimestamp(data _ConnectionData, value uint64) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP, value)
}
func setSettingConnectionReadOnly(data _ConnectionData, value bool) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY, value)
}
func setSettingConnectionZone(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE, value)
}
func setSettingConnectionMaster(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER, value)
}
func setSettingConnectionSlaveType(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE, value)
}
func setSettingConnectionSecondaries(data _ConnectionData, value []string) {
	setConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES, value)
}

// JSON Getter
func getSettingConnectionIdJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID, getSettingConnectionKeyType(NM_SETTING_CONNECTION_ID))
	return
}
func getSettingConnectionUuidJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID, getSettingConnectionKeyType(NM_SETTING_CONNECTION_UUID))
	return
}
func getSettingConnectionTypeJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE, getSettingConnectionKeyType(NM_SETTING_CONNECTION_TYPE))
	return
}
func getSettingConnectionPermissionsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS, getSettingConnectionKeyType(NM_SETTING_CONNECTION_PERMISSIONS))
	return
}
func getSettingConnectionAutoconnectJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT, getSettingConnectionKeyType(NM_SETTING_CONNECTION_AUTOCONNECT))
	return
}
func getSettingConnectionTimestampJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP, getSettingConnectionKeyType(NM_SETTING_CONNECTION_TIMESTAMP))
	return
}
func getSettingConnectionReadOnlyJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY, getSettingConnectionKeyType(NM_SETTING_CONNECTION_READ_ONLY))
	return
}
func getSettingConnectionZoneJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE, getSettingConnectionKeyType(NM_SETTING_CONNECTION_ZONE))
	return
}
func getSettingConnectionMasterJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER, getSettingConnectionKeyType(NM_SETTING_CONNECTION_MASTER))
	return
}
func getSettingConnectionSlaveTypeJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE, getSettingConnectionKeyType(NM_SETTING_CONNECTION_SLAVE_TYPE))
	return
}
func getSettingConnectionSecondariesJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES, getSettingConnectionKeyType(NM_SETTING_CONNECTION_SECONDARIES))
	return
}

// JSON Setter
func setSettingConnectionIdJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_ID))
}
func setSettingConnectionUuidJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_UUID))
}
func setSettingConnectionTypeJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_TYPE))
}
func setSettingConnectionPermissionsJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_PERMISSIONS))
}
func setSettingConnectionAutoconnectJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_AUTOCONNECT))
}
func setSettingConnectionTimestampJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_TIMESTAMP))
}
func setSettingConnectionReadOnlyJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_READ_ONLY))
}
func setSettingConnectionZoneJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_ZONE))
}
func setSettingConnectionMasterJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_MASTER))
}
func setSettingConnectionSlaveTypeJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_SLAVE_TYPE))
}
func setSettingConnectionSecondariesJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES, valueJSON, getSettingConnectionKeyType(NM_SETTING_CONNECTION_SECONDARIES))
}

// Remover
func removeSettingConnectionId(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ID)
}
func removeSettingConnectionUuid(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_UUID)
}
func removeSettingConnectionType(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TYPE)
}
func removeSettingConnectionPermissions(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS)
}
func removeSettingConnectionAutoconnect(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_AUTOCONNECT)
}
func removeSettingConnectionTimestamp(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_TIMESTAMP)
}
func removeSettingConnectionReadOnly(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_READ_ONLY)
}
func removeSettingConnectionZone(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_ZONE)
}
func removeSettingConnectionMaster(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_MASTER)
}
func removeSettingConnectionSlaveType(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SLAVE_TYPE)
}
func removeSettingConnectionSecondaries(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_SECONDARIES)
}
