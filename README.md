# Cisco CallManager API Requests

## AddUserReq

**Purpose**: Represents a request to add a new user to CUCM.

**Structure Fields**:
- `XMLName`: Defines the root element for the XML request.
- `XmlnsSoapenv`: The XML namespace for the SOAP envelope.
- `XmlnsAxlsql`: The XML namespace for the AXL API.
- `Userid`: The user ID for the new user.
- `LastName`: The last name of the new user.
- `FirstName`: The first name of the new user.
- `Password`: The password for the new user.
- `Pin`: The PIN for the new user.
- `TelephoneNumber`: The telephone number for the new user.

**CUCM Function Mapped**: `addUser`  
This function creates a new user in the CUCM database with the specified details.

- **Endpoint**: `/addUser`
- **Method**: `POST`
- **Body**: `JSON`

```
  {
    "userid": "newuser",
    "lastName": "User",
    "firstName": "New",
    "password": "newuserpassword",
    "pin": "12345",
    "telephoneNumber": "1234567890"
  }
```

## AddPhoneReq

**Purpose**: Represents a request to add a new phone to CUCM.

**Structure Fields**:
- `XMLName`: Defines the root element for the XML request.
- `XmlnsSoapenv`: The XML namespace for the SOAP envelope.
- `XmlnsAxlsql`: The XML namespace for the AXL API.
- `Name`: The name (usually the MAC address) of the new phone.
- `Product`: The type of phone (e.g., "Cisco 8841").
- `Class`: The class of the device (e.g., "Phone").
- `Protocol`: The protocol used by the phone (e.g., "SIP").
- `ProtocolSide`: Indicates whether the protocol is user-side or network-side.
- `DevicePoolName`: The name of the device pool the phone belongs to.
- `CommonPhoneConfigName`: The name of the common phone configuration.
- `LocationName`: The name of the location associated with the phone.
- `UseTrustedRelayPoint`: Indicates whether to use a trusted relay point.
- `PhoneTemplateName`: The name of the phone template.
- `Lines`: A structure containing information about the phone lines.

**CUCM Function Mapped**: `addPhone`  
This function adds a new phone to the CUCM database with the specified configuration.

- **Endpoint**: `/addPhone`
- **Method**: `POST`
- **Body**: `JSON`

```
{
  "name": "SEP001122334455",
  "product": "Cisco 8841",
  "class": "Phone",
  "protocol": "SIP",
  "protocolSide": "User",
  "devicePoolName": "Default",
  "commonPhoneConfigName": "Standard Common Phone Profile",
  "locationName": "Hub_None",
  "useTrustedRelayPoint": "Default",
  "phoneTemplateName": "Standard 8841 SIP",
  "lines": [
    {
      "index": 1,
      "pattern": "1001",
      "routePartitionName": "Internal",
      "label": "Line 1"
    }
  ]
}
```
## AssociatePhoneReq

**Purpose**: Represents a request to associate a phone with a user in CUCM.

**Structure Fields**:
- `XMLName`: Defines the root element for the XML request.
- `XmlnsSoapenv`: The XML namespace for the SOAP envelope.
- `XmlnsAxlsql`: The XML namespace for the AXL API.
- `Name`: The name (MAC address) of the phone.
- `OwnerUserName`: The user ID of the phone's owner.
- `Userid`: The user ID to associate with the phone.
- `AssociatedDevices`: A structure containing the list of associated devices.
- `PrimaryExtension`: The primary extension configuration.

**CUCM Function Mapped**: `updatePhone` and `updateUser`  
These functions update the phone to set the owner user ID and associate the device with the user, respectively.

- **Endpoint**: `/associatePhone`
- **Method**: `POST`
- **Body**: `JSON`

```
{
  "name": "SEP001122334455",
  "ownerUserName": "newuser",
  "userid": "newuser",
  "associatedDevices": ["SEP001122334455"],
  "primaryExtension": {
    "pattern": "1001",
    "routePartitionName": "Internal"
  }
}
```

## getUser

- **URL Endpoint**: `/getUser`
- **HTTP Method**: `POST`
- **JSON Request Example**:

```
{
  "userid": "example_user_id"
}
```
## Line

**Purpose**: Represents a single phone line (extension) configuration.

**Structure Fields**:
- `Index`: The index of the line (e.g., 1 for the primary line).
- `Dirn`: The directory number (extension) configuration.
- `Label`: The label for the line.

**CUCM Function Mapped**: Part of the `addPhone` function, specifying the details of each line on the phone.

## Lines (plural)

**Purpose**: Represents the lines (extensions) configuration for a phone.

**Structure Fields**:
- `Line`: A list of Line structures, each representing a phone line.

**CUCM Function Mapped**: Part of the `addPhone` function, configuring the lines for the new phone.

## Dirn

**Purpose**: Represents the directory number (extension) for a phone line.

**Structure Fields**:
- `Pattern`: The extension number.
- `RoutePartitionName`: The route partition to which the extension belongs.

**CUCM Function Mapped**: Part of the `addPhone` function, specifying the extension details.

## Devices

**Purpose**: Represents the list of devices associated with a user.

**Structure Fields**:
- `Device`: A list of device names (e.g., MAC addresses) associated with the user.

**CUCM Function Mapped**: Part of the `updateUser` function, specifying the devices associated with the user.

## Extension

**Purpose**: Represents the primary extension configuration for a user.

**Structure Fields**:
- `Pattern`: The extension number.
- `RoutePartitionName`: The route partition to which the extension belongs.

**CUCM Function Mapped**: Part of the `updateUser` function, specifying the user's primary extension.

## Summary
- **AddUserReq**: Adds a new user to CUCM.
- **AddPhoneReq**: Adds a new phone to CUCM.
- **AssociatePhoneReq**: Associates a phone with a user and sets the primary extension.
