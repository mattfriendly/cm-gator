# AXL API Interface Documentation

This document outlines the AXL API interface provided by our Go microserver for managing Cisco Unified Communications Manager (CUCM) resources. The API allows for the creation and management of users, phones, and associations between them.

## General Information

- **Base URL**: `https://<your-server-address>:8443`
- **Content-Type**: `application/json`
- **Accept**: `application/json`

## Security

- All communications with the API are secured via HTTPS.
- Requests to the API require a valid VPN connection to the CUCM sandbox environment.

## API Endpoints

Each endpoint in the API is designed to manage specific resources within the Cisco Unified Communications Manager. Here, we document the detailed functionality, request formats, and sample calls for the endpoints handling users and phones.

### 1. Add User

- **URL**: `/addUser`
- **Method**: `POST`
- **Description**: Adds a new user to the Cisco Unified Communications Manager.
- **Request Body**:

  ```json
  {
    "userid": "jdoe",
    "lastName": "Doe",
    "firstName": "John",
    "password": "password123",
    "pin": "12345",
    "telephoneNumber": "1001"
  }
  ```

- **Success Response**:

  - **Code**: `200 OK`
  - **Content**:

  ```json
  {
    "status": "success",
    "message": "Operation completed successfully",
    "data": "<response-data-from-CUCM>"
  }
  ```

- **Error Response**:

  - **Code**: `400 Bad Request`
  - **Content**:

  ```json
  {
    "status": "error",
    "message": "Invalid request"
  }
  ```
  - **Code**: `500 Internal Server Error`
  - **Content**:

  ```json
  {
     "status": "error",
     "message": "Failed to forward request"
  }
  ```

  - **Sample Call**:

  ```bash
  curl -X POST https://<your-server-address>:8443/addUser -d '{ "userid": "jdoe", "lastName": "Doe", "firstName": "John", "password": "password123", "pin": "12345", "telephoneNumber": "1001" }' -H "Content-Type: application/json"
  ```


This snippet provides detailed information about handling errors and a sample cURL command for making a request to the `/addUser` endpoint. Let me know if you need further modifications or additional sections!

### 2. Add Phone

- **URL**: `/addPhone`
- **Method**: `POST`
- **Description**: Adds a new user to the Cisco Unified Communications Manager.
- **Request Body**:
    ```json
    {
       "name": "SEP001122334455",
       "product": "Cisco 8841",
       "class": "Phone",
       "protocol": "SIP",
       "protocolSide": "User",
       "devicePoolName": "Default",
       "commonPhoneConfigName": "Standard Common Phone Profile",
       "locationName": "HQ",
       "useTrustedRelayPoint": "Default",
       "phoneTemplateName": "Standard 8841 SIP",
       "lines": {
         "line": [
           {
             "index": 1,
             "dirn": {
               "pattern": "1001",
               "routePartitionName": "Internal"
             },
             "label": "John Doe"
           }
         ]
       }
    }
    ```
- **Success Response**:

  - **Code**: `200 OK`
  - **Content**:

  ```json
  {
    "status": "success",
    "message": "Operation completed successfully",
    "data": "<response-data-from-CUCM>"
  }
  ```

