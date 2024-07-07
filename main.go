package main

import (
        "bytes"
        "crypto/tls"
        "encoding/base64"
        "encoding/json"
        "encoding/xml"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
)

/****
*
* Structures
*
*/

// ExecuteSQLQueryReq structure for SOAP request
type ExecuteSQLQueryReq struct {
        XMLName      xml.Name `xml:"soapenv:Envelope"`
        XmlnsSoapenv string   `xml:"xmlns:soapenv,attr"`
        XmlnsAxl     string   `xml:"xmlns:axl,attr"`
        Body         struct {
                ExecuteSQLQuery struct {
                        SQL string `xml:"sql"`
                } `xml:"axl:executeSQLQuery"`
        } `xml:"soapenv:Body"`
}

// ExecuteSQLQueryResp structure for SOAP response
type ExecuteSQLQueryResp struct {
        XMLName xml.Name `xml:"Envelope"`
        Body    struct {
                ExecuteSQLQueryResponse struct {
                        Return struct {
                                Rows []struct {
                                        UserID     string `xml:"userid"`
                                        FirstName  string `xml:"firstname"`
                                        LastName   string `xml:"lastname"`
                                        Department string `xml:"department"`
                                } `xml:"row"`
                        } `xml:"return"`
                } `xml:"executeSQLQueryResponse"`
        } `xml:"Body"`
}

// Define JsonResponse struct
type JsonResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

// AddPhoneReq structure for JSON request
type AddPhoneReq struct {
        Name                       string  `json:"name"`
        Description                string  `json:"description"`
        Product                    string  `json:"product"`
        Class                      string  `json:"class"`
        Protocol                   string  `json:"protocol"`
        ProtocolSide               string  `json:"protocolSide"`
        CallingSearchSpaceName     string  `json:"callingSearchSpaceName"`
        DevicePoolName             string  `json:"devicePoolName"`
        CommonDeviceConfigName     string  `json:"commonDeviceConfigName"`
        CommonPhoneConfigName      string  `json:"commonPhoneConfigName"`
        NetworkLocation            string  `json:"networkLocation"`
        LocationName               string  `json:"locationName"`
        MediaResourceListName      string  `json:"mediaResourceListName"`
        NetworkHoldMohAudioSourceId string `json:"networkHoldMohAudioSourceId"`
        UserHoldMohAudioSourceId   string  `json:"userHoldMohAudioSourceId"`
        AutomatedAlternateRoutingCssName string `json:"automatedAlternateRoutingCssName"`
        AarNeighborhoodName        string  `json:"aarNeighborhoodName"`
        LoadInformation            struct {
                Special bool   `json:"special"`
                Value   string `json:"value"`
        } `json:"loadInformation"`
        VersionStamp               string `json:"versionStamp"`
        TraceFlag                  bool   `json:"traceFlag"`
        MlppDomainId               string `json:"mlppDomainId"`
        MlppIndicationStatus       string `json:"mlppIndicationStatus"`
        Preemption                 string `json:"preemption"`
        UseTrustedRelayPoint       string `json:"useTrustedRelayPoint"`
        RetryVideoCallAsAudio      bool   `json:"retryVideoCallAsAudio"`
        SecurityProfileName        string `json:"securityProfileName"`
        SipProfileName             string `json:"sipProfileName"`
        CgpnTransformationCssName  string `json:"cgpnTransformationCssName"`
        UseDevicePoolCgpnTransformCss bool `json:"useDevicePoolCgpnTransformCss"`
        GeoLocationName            string `json:"geoLocationName"`
        GeoLocationFilterName      string `json:"geoLocationFilterName"`
        SendGeoLocation            bool   `json:"sendGeoLocation"`
        Lines                      struct {
                Line []struct {
                        Index           int    `json:"index"`
                        Dirn            struct {
                                Pattern            string `json:"pattern"`
                                RoutePartitionName string `json:"routePartitionName"`
                        } `json:"dirn"`
                        Label               string `json:"label"`
                        Display             string `json:"display"`
                        DisplayAscii        string `json:"displayAscii"`
                        E164Mask            string `json:"e164Mask"`
                        DialPlanWizardId    int    `json:"dialPlanWizardId"`
                        MwlPolicy           string `json:"mwlPolicy"`
                        MaxNumCalls         int    `json:"maxNumCalls"`
                        BusyTrigger         int    `json:"busyTrigger"`
                        CallInfoDisplay     struct {
                                CallerName       bool `json:"callerName"`
                                CallerNumber     bool `json:"callerNumber"`
                                RedirectedNumber bool `json:"redirectedNumber"`
                                DialedNumber     bool `json:"dialedNumber"`
                        } `json:"callInfoDisplay"`
                        RecordingProfileName string `json:"recordingProfileName"`
                        MonitoringCssName    string `json:"monitoringCssName"`
                        RecordingFlag        string `json:"recordingFlag"`
                        AudibleMwi           string `json:"audibleMwi"`
                        SpeedDial            string `json:"speedDial"`
                        PartitionUsage       string `json:"partitionUsage"`
                        AssociatedEndusers   struct {
                                Enduser []struct {
                                        UserId string `json:"userId"`
                                } `json:"enduser"`
                        } `json:"associatedEndusers"`
                        MissedCallLogging    bool   `json:"missedCallLogging"`
                        RecordingMediaSource string `json:"recordingMediaSource"`
                } `json:"line"`
        } `json:"lines"`
        NumberOfButtons                 int    `json:"numberOfButtons"`
        PhoneTemplateName               string `json:"phoneTemplateName"`
        Speeddials                      []string `json:"speeddials"`
        BusyLampFields                  []string `json:"busyLampFields"`
        PrimaryPhoneName                string `json:"primaryPhoneName"`
        RingSettingIdleBlfAudibleAlert  string `json:"ringSettingIdleBlfAudibleAlert"`
        RingSettingBusyBlfAudibleAlert  string `json:"ringSettingBusyBlfAudibleAlert"`
        BlfDirectedCallParks            []string `json:"blfDirectedCallParks"`
        AddOnModules                    []string `json:"addOnModules"`
        UserLocale                      string `json:"userLocale"`
        NetworkLocale                   string `json:"networkLocale"`
        IdleTimeout                     int    `json:"idleTimeout"`
        AuthenticationUrl               string `json:"authenticationUrl"`
        DirectoryUrl                    string `json:"directoryUrl"`
        IdleUrl                         string `json:"idleUrl"`
        InformationUrl                  string `json:"informationUrl"`
        MessagesUrl                     string `json:"messagesUrl"`
        ProxyServerUrl                  string `json:"proxyServerUrl"`
        ServicesUrl                     string `json:"servicesUrl"`
        Services                        []string `json:"services"`
        SoftkeyTemplateName             string `json:"softkeyTemplateName"`
        DefaultProfileName              string `json:"defaultProfileName"`
        EnableExtensionMobility         int    `json:"enableExtensionMobility"`
        SingleButtonBarge               string `json:"singleButtonBarge"`
        JoinAcrossLines                 string `json:"joinAcrossLines"`
        BuiltInBridgeStatus             string `json:"builtInBridgeStatus"`
        CallInfoPrivacyStatus           string `json:"callInfoPrivacyStatus"`
        HlogStatus                      string `json:"hlogStatus"`
        OwnerUserName                   string `json:"ownerUserName"`
        IgnorePresentationIndicators    bool   `json:"ignorePresentationIndicators"`
        PacketCaptureMode               string `json:"packetCaptureMode"`
        PacketCaptureDuration           int    `json:"packetCaptureDuration"`
        SubscribeCallingSearchSpaceName string `json:"subscribeCallingSearchSpaceName"`
        RerouteCallingSearchSpaceName   string `json:"rerouteCallingSearchSpaceName"`
        AllowCtiControlFlag             bool   `json:"allowCtiControlFlag"`
        PresenceGroupName               string `json:"presenceGroupName"`
        UnattendedPort                  bool   `json:"unattendedPort"`
        RequireDtmfReception            bool   `json:"requireDtmfReception"`
        Rfc2833Disabled                 bool   `json:"rfc2833Disabled"`
        CertificateOperation            string `json:"certificateOperation"`
        DeviceMobilityMode              string `json:"deviceMobilityMode"`
        RemoteDevice                    bool   `json:"remoteDevice"`
        DndOption                       string `json:"dndOption"`
        DndStatus                       bool   `json:"dndStatus"`
        IsActive                        bool   `json:"isActive"`
        IsDualMode                      bool   `json:"isDualMode"`
        PhoneSuite                      string `json:"phoneSuite"`
        PhoneServiceDisplay             string `json:"phoneServiceDisplay"`
        IsProtected                     bool   `json:"isProtected"`
        MtpRequired                     bool   `json:"mtpRequired"`
        MtpPreferedCodec                string `json:"mtpPreferedCodec"`
        DialRulesName                   string `json:"dialRulesName"`
        SshUserId                       string `json:"sshUserId"`
        DigestUser                      string `json:"digestUser"`
        OutboundCallRollover            string `json:"outboundCallRollover"`
        HotlineDevice                   bool   `json:"hotlineDevice"`
        SecureInformationUrl            string `json:"secureInformationUrl"`
        SecureDirectoryUrl              string `json:"secureDirectoryUrl"`
        SecureMessageUrl                string `json:"secureMessageUrl"`
        SecureServicesUrl               string `json:"secureServicesUrl"`
        SecureAuthenticationUrl         string `json:"secureAuthenticationUrl"`
        SecureIdleUrl                   string `json:"secureIdleUrl"`
        AlwaysUsePrimeLine              bool   `json:"alwaysUsePrimeLine"`
        AlwaysUsePrimeLineForVoiceMessage bool `json:"alwaysUsePrimeLineForVoiceMessage"`
        FeatureControlPolicy            string `json:"featureControlPolicy"`
        DeviceTrustMode                 string `json:"deviceTrustMode"`
        ConfidentialAccess              struct {
                ConfidentialAccessMode string `json:"confidentialAccessMode"`
                ConfidentialAccessLevel string `json:"confidentialAccessLevel"`
        } `json:"confidentialAccess"`
        RequireOffPremiseLocation       bool   `json:"requireOffPremiseLocation"`
        CgpnIngressDN                   string `json:"cgpnIngressDN"`
        UseDevicePoolCgpnIngressDN      bool   `json:"useDevicePoolCgpnIngressDN"`
        Msisdn                          string `json:"msisdn"`
        EnableCallRoutingToRdWhenNoneIsActive bool `json:"enableCallRoutingToRdWhenNoneIsActive"`
        WifiHotspotProfile              string `json:"wifiHotspotProfile"`
        WirelessLanProfileGroup         string `json:"wirelessLanProfileGroup"`
        ElinGroup                       string `json:"elinGroup"`
}

// AddPhoneResp structure for SOAP response
type AddPhoneResp struct {
        XMLName xml.Name `xml:"Envelope"`
        Body    struct {
                AddPhoneResponse struct {
                        Return struct {
                                Name string `xml:"name"`
                        } `xml:"return"`
                } `xml:"addPhoneResponse"`
        } `xml:"Body"`
}

func main() {
        http.HandleFunc("/addPhone", handleAddPhoneRequest)
        http.HandleFunc("/listUsers", handleListUsersRequest)

        // Generate or specify your SSL certificates
        certFile := "./server.crt"
        keyFile := "./server.key"

        log.Println("Starting server on :8443")
        err := http.ListenAndServeTLS(":8443", certFile, keyFile, nil)
        if err != nil {
                log.Fatalf("Server failed to start: %v", err)
        }
}

// Handler function for listing users
func handleListUsersRequest(w http.ResponseWriter, r *http.Request) {
        // Create the SOAP request
        soapRequest := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:axl="http://www.cisco.com/AXL/API/14.0">
   <soapenv:Header/>
   <soapenv:Body>
      <axl:executeSQLQuery>
         <sql>SELECT userid, firstname, lastname, department FROM enduser</sql>
      </axl:executeSQLQuery>
   </soapenv:Body>
</soapenv:Envelope>`

        // Forward the request to Cisco AXL API
        response, err := sendAXLRequest(soapRequest)
        if err != nil {
                http.Error(w, "Failed to forward request", http.StatusInternalServerError)
                logResponse("error", err.Error(), nil)
                return
        }

        // Parse the SOAP response
        var resp ExecuteSQLQueryResp
        if err := xml.Unmarshal(response, &resp); err != nil {
                http.Error(w, "Failed to parse response", http.StatusInternalServerError)
                logResponse("error", err.Error(), nil)
                return
        }

        // Extract user information
        users := make([]map[string]string, len(resp.Body.ExecuteSQLQueryResponse.Return.Rows))
        for i, row := range resp.Body.ExecuteSQLQueryResponse.Return.Rows {
                users[i] = map[string]string{
                        "userid":     row.UserID,
                        "firstname":  row.FirstName,
                        "lastname":   row.LastName,
                        "department": row.Department,
                }
        }

        // Write the JSON response back to the client
        jsonResponse(w, http.StatusOK, "Users retrieved successfully", users)
}

// Handler function for adding a phone
func handleAddPhoneRequest(w http.ResponseWriter, r *http.Request) {
        // Parse the incoming JSON request
        var req AddPhoneReq
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
                http.Error(w, "Invalid request", http.StatusBadRequest)
                logResponse("error", "Invalid request", nil)
                return
        }

        // Manually create the SOAP request
        soapRequest := fmt.Sprintf(`
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:axl="http://www.cisco.com/AXL/API/14.0">
   <soapenv:Header/>
   <soapenv:Body>
      <axl:addPhone>
         <phone>
            <name>%s</name>
            <description>%s</description>
            <product>%s</product>
            <class>%s</class>
            <protocol>%s</protocol>
            <protocolSide>%s</protocolSide>
            <callingSearchSpaceName>%s</callingSearchSpaceName>
            <devicePoolName>%s</devicePoolName>
            <commonDeviceConfigName>%s</commonDeviceConfigName>
            <commonPhoneConfigName>%s</commonPhoneConfigName>
            <networkLocation>%s</networkLocation>
            <locationName>%s</locationName>
            <mediaResourceListName>%s</mediaResourceListName>
            <networkHoldMohAudioSourceId>%s</networkHoldMohAudioSourceId>
            <userHoldMohAudioSourceId>%s</userHoldMohAudioSourceId>
            <automatedAlternateRoutingCssName>%s</automatedAlternateRoutingCssName>
            <aarNeighborhoodName>%s</aarNeighborhoodName>
            <loadInformation special="%t">%s</loadInformation>
            <vendorConfig/>
            <versionStamp>%s</versionStamp>
            <traceFlag>%t</traceFlag>
            <mlppDomainId>%s</mlppDomainId>
            <mlppIndicationStatus>%s</mlppIndicationStatus>
            <preemption>%s</preemption>
            <useTrustedRelayPoint>%s</useTrustedRelayPoint>
            <retryVideoCallAsAudio>%t</retryVideoCallAsAudio>
            <securityProfileName>%s</securityProfileName>
            <sipProfileName>%s</sipProfileName>
            <cgpnTransformationCssName>%s</cgpnTransformationCssName>
            <useDevicePoolCgpnTransformCss>%t</useDevicePoolCgpnTransformCss>
            <geoLocationName>%s</geoLocationName>
            <geoLocationFilterName>%s</geoLocationFilterName>
            <sendGeoLocation>%t</sendGeoLocation>
            <lines>
               <line>
                  <index>%d</index>
                  <dirn>
                     <pattern>%s</pattern>
                     <routePartitionName>%s</routePartitionName>
                  </dirn>
                  <label>%s</label>
                  <display>%s</display>
                  <displayAscii>%s</displayAscii>
                  <e164Mask>%s</e164Mask>
                  <dialPlanWizardId>%d</dialPlanWizardId>
                  <mwlPolicy>%s</mwlPolicy>
                  <maxNumCalls>%d</maxNumCalls>
                  <busyTrigger>%d</busyTrigger>
                  <callInfoDisplay>
                     <callerName>%t</callerName>
                     <callerNumber>%t</callerNumber>
                     <redirectedNumber>%t</redirectedNumber>
                     <dialedNumber>%t</dialedNumber>
                  </callInfoDisplay>
                  <recordingProfileName>%s</recordingProfileName>
                  <monitoringCssName>%s</monitoringCssName>
                  <recordingFlag>%s</recordingFlag>
                  <audibleMwi>%s</audibleMwi>
                  <speedDial>%s</speedDial>
                  <partitionUsage>%s</partitionUsage>
                  <associatedEndusers>
                     <enduser>
                        <userId>%s</userId>
                     </enduser>
                  </associatedEndusers>
                  <missedCallLogging>%t</missedCallLogging>
                  <recordingMediaSource>%s</recordingMediaSource>
               </line>
            </lines>
            <numberOfButtons>%d</numberOfButtons>
            <phoneTemplateName>%s</phoneTemplateName>
            <speeddials>%s</speeddials>
            <busyLampFields>%s</busyLampFields>
            <primaryPhoneName>%s</primaryPhoneName>
            <ringSettingIdleBlfAudibleAlert>%s</ringSettingIdleBlfAudibleAlert>
            <ringSettingBusyBlfAudibleAlert>%s</ringSettingBusyBlfAudibleAlert>
            <blfDirectedCallParks>%s</blfDirectedCallParks>
            <addOnModules>%s</addOnModules>
            <userLocale>%s</userLocale>
            <networkLocale>%s</networkLocale>
            <idleTimeout>%d</idleTimeout>
            <authenticationUrl>%s</authenticationUrl>
            <directoryUrl>%s</directoryUrl>
            <idleUrl>%s</idleUrl>
            <informationUrl>%s</informationUrl>
            <messagesUrl>%s</messagesUrl>
            <proxyServerUrl>%s</proxyServerUrl>
            <servicesUrl>%s</servicesUrl>
            <services>%s</services>
            <softkeyTemplateName>%s</softkeyTemplateName>
            <defaultProfileName>%s</defaultProfileName>
            <enableExtensionMobility>%d</enableExtensionMobility>
            <singleButtonBarge>%s</singleButtonBarge>
            <joinAcrossLines>%s</joinAcrossLines>
            <builtInBridgeStatus>%s</builtInBridgeStatus>
            <callInfoPrivacyStatus>%s</callInfoPrivacyStatus>
            <hlogStatus>%s</hlogStatus>
            <ownerUserName>%s</ownerUserName>
            <ignorePresentationIndicators>%t</ignorePresentationIndicators>
            <packetCaptureMode>%s</packetCaptureMode>
            <packetCaptureDuration>%d</packetCaptureDuration>
            <subscribeCallingSearchSpaceName>%s</subscribeCallingSearchSpaceName>
            <rerouteCallingSearchSpaceName>%s</rerouteCallingSearchSpaceName>
            <allowCtiControlFlag>%t</allowCtiControlFlag>
            <presenceGroupName>%s</presenceGroupName>
            <unattendedPort>%t</unattendedPort>
            <requireDtmfReception>%t</requireDtmfReception>
            <rfc2833Disabled>%t</rfc2833Disabled>
            <certificateOperation>%s</certificateOperation>
            <deviceMobilityMode>%s</deviceMobilityMode>
            <remoteDevice>%t</remoteDevice>
            <dndOption>%s</dndOption>
            <dndStatus>%t</dndStatus>
            <isActive>%t</isActive>
            <isDualMode>%t</isDualMode>
            <phoneSuite>%s</phoneSuite>
            <phoneServiceDisplay>%s</phoneServiceDisplay>
            <isProtected>%t</isProtected>
            <mtpRequired>%t</mtpRequired>
            <mtpPreferedCodec>%s</mtpPreferedCodec>
            <dialRulesName>%s</dialRulesName>
            <sshUserId>%s</sshUserId>
            <digestUser>%s</digestUser>
            <outboundCallRollover>%s</outboundCallRollover>
            <hotlineDevice>%t</hotlineDevice>
            <secureInformationUrl>%s</secureInformationUrl>
            <secureDirectoryUrl>%s</secureDirectoryUrl>
            <secureMessageUrl>%s</secureMessageUrl>
            <secureServicesUrl>%s</secureServicesUrl>
            <secureAuthenticationUrl>%s</secureAuthenticationUrl>
            <secureIdleUrl>%s</secureIdleUrl>
            <alwaysUsePrimeLine>%t</alwaysUsePrimeLine>
            <alwaysUsePrimeLineForVoiceMessage>%t</alwaysUsePrimeLineForVoiceMessage>
            <featureControlPolicy>%s</featureControlPolicy>
            <deviceTrustMode>%s</deviceTrustMode>
            <confidentialAccess>
               <confidentialAccessMode>%s</confidentialAccessMode>
               <confidentialAccessLevel>%s</confidentialAccessLevel>
            </confidentialAccess>
            <requireOffPremiseLocation>%t</requireOffPremiseLocation>
            <cgpnIngressDN>%s</cgpnIngressDN>
            <useDevicePoolCgpnIngressDN>%t</useDevicePoolCgpnIngressDN>
            <msisdn>%s</msisdn>
            <enableCallRoutingToRdWhenNoneIsActive>%t</enableCallRoutingToRdWhenNoneIsActive>
            <wifiHotspotProfile>%s</wifiHotspotProfile>
            <wirelessLanProfileGroup>%s</wirelessLanProfileGroup>
            <elinGroup>%s</elinGroup>
         </phone>
      </axl:addPhone>
   </soapenv:Body>
</soapenv:Envelope>`,
                req.Name,
                req.Description,
                req.Product,
                req.Class,
                req.Protocol,
                req.ProtocolSide,
                req.CallingSearchSpaceName,
                req.DevicePoolName,
                req.CommonDeviceConfigName,
                req.CommonPhoneConfigName,
                req.NetworkLocation,
                req.LocationName,
                req.MediaResourceListName,
                req.NetworkHoldMohAudioSourceId,
                req.UserHoldMohAudioSourceId,
                req.AutomatedAlternateRoutingCssName,
                req.AarNeighborhoodName,
                req.LoadInformation.Special,
                req.LoadInformation.Value,
                req.VersionStamp,
                req.TraceFlag,
                req.MlppDomainId,
                req.MlppIndicationStatus,
                req.Preemption,
                req.UseTrustedRelayPoint,
                req.RetryVideoCallAsAudio,
                req.SecurityProfileName,
                req.SipProfileName,
                req.CgpnTransformationCssName,
                req.UseDevicePoolCgpnTransformCss,
                req.GeoLocationName,
                req.GeoLocationFilterName,
                req.SendGeoLocation,
                req.Lines.Line[0].Index,
                req.Lines.Line[0].Dirn.Pattern,
                req.Lines.Line[0].Dirn.RoutePartitionName,
                req.Lines.Line[0].Label,
                req.Lines.Line[0].Display,
                req.Lines.Line[0].DisplayAscii,
                req.Lines.Line[0].E164Mask,
                req.Lines.Line[0].DialPlanWizardId,
                req.Lines.Line[0].MwlPolicy,
                req.Lines.Line[0].MaxNumCalls,
                req.Lines.Line[0].BusyTrigger,
                req.Lines.Line[0].CallInfoDisplay.CallerName,
                req.Lines.Line[0].CallInfoDisplay.CallerNumber,
                req.Lines.Line[0].CallInfoDisplay.RedirectedNumber,
                req.Lines.Line[0].CallInfoDisplay.DialedNumber,
                req.Lines.Line[0].RecordingProfileName,
                req.Lines.Line[0].MonitoringCssName,
                req.Lines.Line[0].RecordingFlag,
                req.Lines.Line[0].AudibleMwi,
                req.Lines.Line[0].SpeedDial,
                req.Lines.Line[0].PartitionUsage,
                req.Lines.Line[0].AssociatedEndusers.Enduser[0].UserId,
                req.Lines.Line[0].MissedCallLogging,
                req.Lines.Line[0].RecordingMediaSource,
                req.NumberOfButtons,
                req.PhoneTemplateName,
                "speeddials_placeholder",
                "busyLampFields_placeholder",
                req.PrimaryPhoneName,
                req.RingSettingIdleBlfAudibleAlert,
                req.RingSettingBusyBlfAudibleAlert,
                "blfDirectedCallParks_placeholder",
                "addOnModules_placeholder",
                req.UserLocale,
                req.NetworkLocale,
                req.IdleTimeout,
                req.AuthenticationUrl,
                req.DirectoryUrl,
                req.IdleUrl,
                req.InformationUrl,
                req.MessagesUrl,
                req.ProxyServerUrl,
                req.ServicesUrl,
                "services_placeholder",
                req.SoftkeyTemplateName,
                req.DefaultProfileName,
                req.EnableExtensionMobility,
                req.SingleButtonBarge,
                req.JoinAcrossLines,
                req.BuiltInBridgeStatus,
                req.CallInfoPrivacyStatus,
                req.HlogStatus,
                req.OwnerUserName,
                req.IgnorePresentationIndicators,
                req.PacketCaptureMode,
                req.PacketCaptureDuration,
                req.SubscribeCallingSearchSpaceName,
                req.RerouteCallingSearchSpaceName,
                req.AllowCtiControlFlag,
                req.PresenceGroupName,
                req.UnattendedPort,
                req.RequireDtmfReception,
                req.Rfc2833Disabled,
                req.CertificateOperation,
                req.DeviceMobilityMode,
                req.RemoteDevice,
                req.DndOption,
                req.DndStatus,
                req.IsActive,
                req.IsDualMode,
                req.PhoneSuite,
                req.PhoneServiceDisplay,
                req.IsProtected,
                req.MtpRequired,
                req.MtpPreferedCodec,
                req.DialRulesName,
                req.SshUserId,
                req.DigestUser,
                req.OutboundCallRollover,
                req.HotlineDevice,
                req.SecureInformationUrl,
                req.SecureDirectoryUrl,
                req.SecureMessageUrl,
                req.SecureServicesUrl,
                req.SecureAuthenticationUrl,
                req.SecureIdleUrl,
                req.AlwaysUsePrimeLine,
                req.AlwaysUsePrimeLineForVoiceMessage,
                req.FeatureControlPolicy,
                req.DeviceTrustMode,
                req.ConfidentialAccess.ConfidentialAccessMode,
                req.ConfidentialAccess.ConfidentialAccessLevel,
                req.RequireOffPremiseLocation,
                req.CgpnIngressDN,
                req.UseDevicePoolCgpnIngressDN,
                req.Msisdn,
                req.EnableCallRoutingToRdWhenNoneIsActive,
                req.WifiHotspotProfile,
                req.WirelessLanProfileGroup,
                req.ElinGroup)

        // Forward the request to Cisco AXL API
        response, err := sendAXLRequest(soapRequest)
        if err != nil {
                http.Error(w, "Failed to forward request", http.StatusInternalServerError)
                logResponse("error", err.Error(), nil)
                return
        }

        // Parse the SOAP response
        var resp AddPhoneResp
        if err := xml.Unmarshal(response, &resp); err != nil {
                http.Error(w, "Failed to parse response", http.StatusInternalServerError)
                logResponse("error", err.Error(), nil)
                return
        }

        // Write the JSON response back to the client
        jsonResponse(w, http.StatusOK, "Phone added successfully", resp.Body.AddPhoneResponse.Return)
}

/****
*
* Helper functions
*
*/

// Function to send AXL requests
func sendAXLRequest(soapRequest string) ([]byte, error) {
        // Set up the HTTP client with TLS configuration
        httpClient := &http.Client{
                Transport: &http.Transport{
                        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
                },
        }

        // Create the HTTP request
        req, err := http.NewRequest("POST", "https://10.10.20.1:8443/axl/", bytes.NewBuffer([]byte(soapRequest)))
        if err != nil {
                return nil, fmt.Errorf("failed to create HTTP request: %v", err)
        }
        req.Header.Set("Content-Type", "text/xml")
        req.Header.Set("SOAPAction", "CUCM:DB ver=14.0")

        // Add Basic Authentication header
        username := "<redacted>"
        password := "<redacted>"
        auth := username + ":" + password
        req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

        // Send the HTTP request
        resp, err := httpClient.Do(req)
        if err != nil {
                return nil, fmt.Errorf("failed to send HTTP request: %v", err)
        }
        defer resp.Body.Close()

        // Read the response body
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return nil, fmt.Errorf("failed to read response body: %v", err)
        }

        // Log the HTML response for debugging
        if resp.Header.Get("Content-Type") == "text/html" {
                log.Printf("Received HTML response: %s", string(body))
        }

        return body, nil
}


// Function to send JSON responses
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

// Function to log responses
func logResponse(status, message string, data interface{}) {
        logData := JsonResponse{
                Status:  status,
                Message: message,
                Data:    data,
        }
        log.Println(logData)
}
