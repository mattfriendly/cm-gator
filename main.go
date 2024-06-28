package main

/****
*
* Import section
*
 */

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/tiaguinho/gosoap"
)

/****
*
* Structures
*
 */

// AddUserReq structure for SOAP request
type AddUserReq struct {
	XMLName         xml.Name `xml:"soapenv:Envelope"`
	XmlnsSoapenv    string   `xml:"xmlns:soapenv,attr"`
	XmlnsAxlsql     string   `xml:"xmlns:axl,attr"`
	Userid          string   `xml:"axl:userid"`
	LastName        string   `xml:"axl:lastName"`
	FirstName       string   `xml:"axl:firstName"`
	Password        string   `xml:"axl:password"`
	Pin             string   `xml:"axl:pin"`
	TelephoneNumber string   `xml:"axl:telephoneNumber"`
}

// AddPhoneReq structure for SOAP request
type AddPhoneReq struct {
	XMLName               xml.Name `xml:"soapenv:Envelope"`
	XmlnsSoapenv          string   `xml:"xmlns:soapenv,attr"`
	XmlnsAxlsql           string   `xml:"xmlns:axl,attr"`
	Name                  string   `xml:"axl:name"`
	Product               string   `xml:"axl:product"`
	Class                 string   `xml:"axl:class"`
	Protocol              string   `xml:"axl:protocol"`
	ProtocolSide          string   `xml:"axl:protocolSide"`
	DevicePoolName        string   `xml:"axl:devicePoolName"`
	CommonPhoneConfigName string   `xml:"axl:commonPhoneConfigName"`
	LocationName          string   `xml:"axl:locationName"`
	UseTrustedRelayPoint  string   `xml:"axl:useTrustedRelayPoint"`
	PhoneTemplateName     string   `xml:"axl:phoneTemplateName"`
	Lines                 Lines    `xml:"axl:lines"`
}

// Lines structure for SOAP request
type Lines struct {
	Line []Line `xml:"axl:line"`
}

// Line structure for SOAP request
type Line struct {
	Index int    `xml:"axl:index"`
	Dirn  Dirn   `xml:"axl:dirn"`
	Label string `xml:"axl:label"`
}

// Dirn structure for SOAP request
type Dirn struct {
	Pattern            string `xml:"axl:pattern"`
	RoutePartitionName string `xml:"axl:routePartitionName"`
}

// AssociatePhoneReq structure for SOAP request
type AssociatePhoneReq struct {
	XMLName           xml.Name  `xml:"soapenv:Envelope"`
	XmlnsSoapenv      string    `xml:"xmlns:soapenv,attr"`
	XmlnsAxlsql       string    `xml:"xmlns:axl,attr"`
	Name              string    `xml:"axl:name"`
	OwnerUserName     string    `xml:"axl:ownerUserName"`
	Userid            string    `xml:"axl:userid"`
	AssociatedDevices Devices   `xml:"axl:associatedDevices"`
	PrimaryExtension  Extension `xml:"axl:primaryExtension"`
}

// Devices structure for SOAP request
type Devices struct {
	Device []string `xml:"axl:device"`
}

// Extension structure for SOAP request
type Extension struct {
	Pattern            string `xml:"axl:pattern"`
	RoutePartitionName string `xml:"axl:routePartitionName"`
}

// General JSON response structure
type JsonResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// GetUserReq structure for SOAP request
type GetUserReq struct {
	XMLName      xml.Name `xml:"soapenv:Envelope"`
	XmlnsSoapenv string   `xml:"xmlns:soapenv,attr"`
	XmlnsAxl     string   `xml:"xmlns:axl,attr"`
	UserID       string   `xml:"axl:userid"`
}

func main() {
	// Serving files from the static files directory
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// API endpoints
	http.HandleFunc("/main", handleMainPage)
	http.HandleFunc("/addUser", handleAddUserRequest)
	http.HandleFunc("/addPhone", handleAddPhoneRequest)
	http.HandleFunc("/associatePhone", handleAssociatePhoneRequest)
	http.HandleFunc("/getUser", handleGetUserRequest)

	// a test endpoint
	http.HandleFunc("/api.data", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		// Use id for something, or remove if not needed
		fmt.Fprintf(w, "ID: %s", id) // Example response using id
	})

	// Generate or specify your SSL certificates
	certFile := "./cert.pem"
	keyFile := "./key.pem"

	log.Println("Starting server on https://localhost:8443")
	err := http.ListenAndServeTLS(":8443", certFile, keyFile, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

/****
*
* Request handler functions
*
 */

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func handleAddUserRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request
	var req AddUserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		logResponse("error", "Invalid request", nil)
		return
	}

	// Forward the request to Cisco AXL API
	response, err := sendAXLRequest(req, "addUser")
	if err != nil {
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		logResponse("error", err.Error(), nil)
		return
	}

	// Write the JSON response back to the client
	jsonResponse(w, http.StatusOK, "Operation completed successfully", response)
}

func handleAddPhoneRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request
	var req AddPhoneReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		logResponse("error", "Invalid request", nil)
		return
	}

	// Forward the request to Cisco AXL API
	response, err := sendAXLRequest(req, "addPhone")
	if err != nil {
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		logResponse("error", err.Error(), nil)
		return
	}

	// Write the JSON response back to the client
	jsonResponse(w, http.StatusOK, "Operation completed successfully", response)
}

func handleAssociatePhoneRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request
	var req AssociatePhoneReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		logResponse("error", "Invalid request", nil)
		return
	}

	// Forward the request to Cisco AXL API
	response, err := sendAXLRequest(req, "updatePhone")
	if err != nil {
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		logResponse("error", err.Error(), nil)
		return
	}

	// Write the JSON response back to the client
	jsonResponse(w, http.StatusOK, "Operation completed successfully", response)
}

func handleGetUserRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request to get the UserID
	var req struct {
		UserID string `json:"userid"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Create the SOAP request
	getUserReq := GetUserReq{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsAxl:     "http://www.cisco.com/AXL/API/14.0",
		UserID:       req.UserID,
	}

	// Forward the request to Cisco AXL API
	response, err := sendAXLRequest(getUserReq, "getUser")
	if err != nil {
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		return
	}

	// Write the JSON response back to the client
	jsonResponse(w, http.StatusOK, "User retrieved successfully", response)
}

/****
*
* AXL request functions
*
 */

func sendAXLRequest(req interface{}, method string) (interface{}, error) {
	// Set up the HTTP client with TLS configuration
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	client, err := gosoap.SoapClient("https://<CUCM_ADDRESS>:8443/axl/", httpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %v", err)
	}

	resp, err := client.Call(method, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call AXL API: %v", err)
	}

	return resp, nil
}

func jsonResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := JsonResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
	logResponse("success", message, data)
}

func logResponse(status, message string, data interface{}) {
	logData := JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	log.Println(logData)
}
