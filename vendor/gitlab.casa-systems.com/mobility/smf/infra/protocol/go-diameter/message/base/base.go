package base

//----------------------- CASA enhancement --------------------------

var BaseDictionary = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

	<application id="0" name="Base"> <!-- Diameter Common Messages -->

		<command code="257" short="CE" name="Capabilities-Exchange">
			<request>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Host-IP-Address" required="true" min="1"/>
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="Product-Name" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Supported-Vendor-Id" required="False"/>
				<rule avp="Auth-Application-Id" required="False"/>
				<rule avp="Acct-Application-Id" required="False"/>
				<rule avp="Vendor-Specific-Application-Id" required="False"/>
				<rule avp="Firmware-Revision" required="False" max="1"/>
			</request>
			<answer>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Host-IP-Address" required="true" min="1"/>
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="Product-Name" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Supported-Vendor-Id" required="False"/>
				<rule avp="Auth-Application-Id" required="False"/>
				<rule avp="Acct-Application-Id" required="False"/>
				<rule avp="Vendor-Specific-Application-Id" required="False"/>
				<rule avp="Firmware-Revision" required="False" max="1"/>
			</answer>
		</command>

		<command code="258" short="RA" name="Re-Auth">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Destination-Host" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Re-Auth-Request-Type" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="265" short="AA" name="AA">
			<request>
				<!-- 9.2.2.2.1	AA-Request (AAR) Command & 9.2.2.5.1-->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="DRMP" required="false" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Auth-Request-Type" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="MIP6-Agent-Info" required="false" max="1"/>
				<rule avp="MIP6-Feature-Vector" required="false" max="1"/>
				<rule avp="Visited-Network-Identifier" required="false" max="1"/>
				<rule avp="QoS-Capability" required="false" max="1"/>
				<rule avp="Service-Selection" required="false" max="1"/>
				<rule avp="OC-Supported-Features" required="false" max="1"/>
				<rule avp="Origination-Time-Stamp" required="false" max="1"/>
				<rule avp="Maximum-Wait-Time" required="false" max="1"/>
				<rule avp="Supported-Features" required="false"/>
				<rule avp="Emergency-Services" required="false" max="1"/>
			</request>
			<answer>
				<!-- 9.2.2.2.2	AA-Answer (AAA) Command & 9.2.2.5.2 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="DRMP" required="false" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Auth-Request-Type" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="MIP6-Feature-Vector" required="false" max="1"/>
				<rule avp="Session-Timeout" required="false" max="1"/>
				<rule avp="APN-Configuration" required="false" max="1"/>
				<rule avp="QoS-Resources" required="false" max="1"/>
				<rule avp="AN-Trusted" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Trace-Info" required="false" max="1"/>
				<rule avp="OC-Supported-Features" required="false" max="1"/>
				<rule avp="OC-OLR" required="false" max="1"/>
				<rule avp="Load" required="false"/>
				<rule avp="Supported-Features" required="false"/>
				<rule avp="MIP-Session-Key" required="false" max="1"/>
			</answer>
		</command>

		<command code="271" short="AC" name="Accounting">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Accounting-Record-Type" required="true" max="1"/>
				<rule avp="Accounting-Record-Number" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="false" max="1"/>
				<rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="Accounting-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Interim-Interval" required="false" max="1"/>
				<rule avp="Accounting-Realtime-Required" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Accounting-Record-Type" required="true" max="1"/>
				<rule avp="Accounting-Record-Number" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="false" max="1"/>
				<rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Accounting-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Acct-Interim-Interval" required="false" max="1"/>
				<rule avp="Accounting-Realtime-Required" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="274" short="AS" name="Abort-Session">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Destination-Host" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="275" short="ST" name="Session-Termination">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Termination-Cause" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="280" short="DW" name="Device-Watchdog">
			<request>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
			</request>
			<answer>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
			</answer>
		</command>

		<command code="282" short="DP" name="Disconnect-Peer">
			<request>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Disconnect-Cause" required="false" max="1"/>
			</request>
			<answer>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
			</answer>
		</command>

		<command code="272" short="CC" name="Credit-Control">
			<request>
				<!-- http://tools.ietf.org/html/rfc4006#section-3.1 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Service-Context-Id" required="true" max="1"/>
				<rule avp="CC-Request-Type" required="true" max="1"/>
				<rule avp="CC-Request-Number" required="true" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="CC-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Subscription-Id" required="false" max="1"/>
				<rule avp="Service-Identifier" required="false" max="1"/>
				<rule avp="Termination-Cause" required="false" max="1"/>
				<rule avp="Requested-Service-Unit" required="false" max="1"/>
				<rule avp="Requested-Action" required="false" max="1"/>
				<rule avp="Used-Service-Unit" required="false" max="1"/>
				<rule avp="Multiple-Services-Indicator" required="false" max="1"/>
				<rule avp="Multiple-Services-Credit-Control" required="false" max="1"/>
				<rule avp="Service-Parameter-Info" required="false" max="1"/>
				<rule avp="CC-Correlation-Id" required="false" max="1"/>
				<rule avp="User-Equipment-Info" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false" max="1"/>
				<rule avp="Route-Record" required="false" max="1"/>
				<rule avp="Service-Information" required="false" max="1"/>
			</request>
			<answer>
				<!-- http://tools.ietf.org/html/rfc4006#section-3.2 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="CC-Request-Type" required="true" max="1"/>
				<rule avp="CC-Request-Number" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="CC-Session-Failover" required="false" max="1"/>
				<rule avp="CC-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Granted-Service-Unit" required="false" max="1"/>
				<rule avp="Multiple-Services-Credit-Control" required="false" max="1"/>
				<rule avp="Cost-Information" required="false" max="1"/>
				<rule avp="Final-Unit-Indication" required="false" max="1"/>
				<rule avp="Check-Balance-Result" required="false" max="1"/>
				<rule avp="Credit-Control-Failure-Handling" required="false" max="1"/>
				<rule avp="Direct-Debiting-Failure-Handling" required="false" max="1"/>
				<rule avp="Validity-Time" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false" max="1"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false" max="1"/>
				<rule avp="Route-Record" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
			</answer>
		</command>

		<avp name="Acct-Interim-Interval" code="85" must="M" may="-" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Accounting-Realtime-Required" code="483" must="-" may="-" must-not="-" may-encrypt="Y">
			<data type="Enumerated">
				<item code="1" name="DELIVER_AND_GRANT"/>
				<item code="2" name="GRANT_AND_STORE"/>
				<item code="3" name="GRANT_AND_LOSE"/>
			</data>
		</avp>

		<avp name="Acct-Multi-Session-Id" code="50" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

		<avp name="Accounting-Record-Number" code="485" must="M" may="-" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Accounting-Record-Type" code="480" must="M" may="-" must-not="V" may-encrypt="Y">
			<data type="Enumerated">
				<item code="1" name="EVENT_RECORD"/>
				<item code="2" name="START_RECORD"/>
				<item code="3" name="INTERIM_RECORD"/>
				<item code="4" name="STOP_RECORD"/>
			</data>
		</avp>

		<avp name="Accounting-Session-Id" code="44" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="OctetString"/>
		</avp>

		<avp name="Accounting-Sub-Session-Id" code="287" must="-" may="-" must-not="-" may-encrypt="Y">
			<data type="Unsigned64"/>
		</avp>

		<avp name="Acct-Application-Id" code="259" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Application-Id" code="258" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Request-Type" code="274" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="1" name="AUTHENTICATE_ONLY"/>
				<item code="2" name="AUTHORIZE_ONLY"/>
				<item code="3" name="AUTHORIZE_AUTHENTICATE"/>
			</data>
		</avp>

		<avp name="Authorization-Lifetime" code="291" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Grace-Period" code="276" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Session-State" code="277" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="STATE_MAINTAINED"/>
				<item code="1" name="NO_STATE_MAINTAINED"/>
			</data>
		</avp>

		<avp name="Re-Auth-Request-Type" code="285" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="AUTHORIZE_ONLY"/>
				<item code="1" name="AUTHORIZE_AUTHENTICATE"/>
			</data>
		</avp>

		<avp name="Class" code="25" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="OctetString"/>
		</avp>

		<avp name="Destination-Host" code="293" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Destination-Realm" code="283" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Disconnect-Cause" code="273" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="REBOOTING"/>
				<item code="1" name="BUSY"/>
				<item code="2" name="DO_NOT_WANT_TO_TALK_TO_YOU"/>
			</data>
		</avp>

		<avp name="Error-Message" code="281" must="-" may="-" must-not="-" may-encrypt="-">
			<data type="UTF8String"/>
		</avp>

		<avp name="Error-Reporting-Host" code="294" must="-" may="-" must-not="V,M" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Event-Timestamp" code="55" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Time"/>
		</avp>

		<avp name="Experimental-Result" code="297" must="-" may="-" must-not="V,M" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="Experimental-Result-Code" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Experimental-Result-Code" code="298" must="-" may="-" must-not="V,M" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Failed-AVP" code="279" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Grouped"/>
		</avp>

		<avp name="Firmware-Revision" code="267" must="-" may="-" must-not="P,V,M" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Host-IP-Address" code="257" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Address"/>
		</avp>

		<avp name="Inband-Security-Id" code="299" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Multi-Round-Time-Out" code="272" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Origin-Host" code="264" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Origin-Realm" code="296" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Product-Name" code="269" must="-" may="-" must-not="P,V,M" may-encrypt="-">
			<data type="UTF8String"/>
		</avp>

		<avp name="Proxy-Host" code="280" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Proxy-Info" code="284" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Proxy-Host" required="true" max="1"/>
				<rule avp="Proxy-State" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Proxy-State" code="33" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="OctetString"/>
		</avp>

		<avp name="Redirect-Host" code="292" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="DiameterURI"/>
		</avp>

		<avp name="Redirect-Host-Usage" code="261" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="DONT_CACHE"/>
				<item code="1" name="ALL_SESSION"/>
				<item code="2" name="ALL_REALM"/>
				<item code="3" name="REALM_AND_APPLICATION"/>
				<item code="4" name="ALL_APPLICATION"/>
				<item code="5" name="ALL_HOST"/>
				<item code="6" name="ALL_USER"/>
			</data>
		</avp>

		<avp name="Redirect-Max-Cache-Time" code="262" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Result-Code" code="268" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Route-Record" code="282" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Session-Id" code="263" must="M" may="-" must-not="V" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

		<avp name="Session-Timeout" code="27" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Session-Binding" code="270" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Session-Server-Failover" code="271" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Enumerated">
				<item code="0" name="REFUSE_SERVICE"/>
				<item code="1" name="TRY_AGAIN"/>
				<item code="2" name="ALLOW_SERVICE"/>
				<item code="3" name="TRY_AGAIN_ALLOW_SERVICE"/>
			</data>
		</avp>

		<avp name="Supported-Vendor-Id" code="265" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Termination-Cause" code="295" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="1" name="DIAMETER_LOGOUT"/>
				<item code="2" name="DIAMETER_SERVICE_NOT_PROVIDED"/>
				<item code="3" name="DIAMETER_BAD_ANSWER"/>
				<item code="4" name="DIAMETER_ADMINISTRATIVE"/>
				<item code="5" name="DIAMETER_LINK_BROKEN"/>
				<item code="6" name="DIAMETER_AUTH_EXPIRED"/>
				<item code="7" name="DIAMETER_USER_MOVED"/>
				<item code="8" name="DIAMETER_SESSION_TIMEOUT"/>
			</data>
		</avp>

		<avp name="User-Name" code="1" must="M" may="-" must-not="V" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

		<avp name="Vendor-Id" code="266" must="-" may="-" must-not="-" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Vendor-Specific-Application-Id" code="260" must="-" may="-" must-not="-" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Vendor-Id" required="false" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="true" max="1"/>
			</data>
		</avp>

		<!-- IETF RFC 7683 - https://tools.ietf.org/html/rfc7683 -->
		<avp name="OC-Supported-Features" code="621" must-not="V,M">
			<data type="Grouped">
				<rule avp="OC-Feature-Vector" required="false"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="OC-Feature-Vector" code="622" must-not="V,M">
			<data type="Unsigned64"/>
		</avp>

		<avp name="OC-OLR" code="623" must-not="V,M">
			<data type="Grouped">
				<rule avp="OC-Sequence-Number" required="true" max="1"/>
				<rule avp="OC-Report-Type" required="true" max="1"/>
				<rule avp="OC-Reduction-Percentage" required="false" max="1"/>
				<rule avp="OC-Validity-Duration" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="OC-Sequence-Number" code="624" must-not="V,M">
			<data type="Unsigned64"/>
		</avp>

		<avp name="OC-Validity-Duration" code="625" must-not="V,M">
			<data type="Unsigned32"/>
		</avp>

		<avp name="OC-Report-Type" code="626" must-not="V,M">
			<data type="Enumerated">
				<item code="0" name="HOST_REPORT"/>
				<item code="1" name="REALM_REPORT"/>
			</data>
		</avp>

		<avp name="OC-Reduction-Percentage" code="627" must-not="V,M">
			<data type="Unsigned32"/>
		</avp>
		
		<avp name="Tariff-Time-Change" code="451" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.20-->
			<data type="Time"/>
		</avp>

		<avp name="CC-Time" code="420" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.21 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="CC-Money" code="413" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.22 -->
			<data type="Grouped">
				<rule avp="Unit-Value" required="true" max="1"/>
				<rule avp="Currency-Code" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Unit-Value" code="445" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.8-->
			<data type="Grouped">
				<rule avp="Value-Digits" required="true" max="1"/>
				<rule avp="Exponent" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Exponent" code="429" must="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.9 -->
			<data type="Integer32"/>
		</avp>

		<avp name="Currency-Code" code="425" must="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.11 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="CC-Total-Octets" code="421" must="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.23 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Input-Octets" code="412" must="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.24 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Output-Octets" code="414" must="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.25 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Service-Specific-Units" code="417" must="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.26 -->
			<data type="Unsigned64"/>
		</avp>

		<!-- IETF RFC 7944 - https://tools.ietf.org/html/rfc7944 -->
		<avp name="DRMP" code="301" must-not="M,V">
			<data type="Enumerated">
				<item code="0" name="PRIORITY_0"/>
				<item code="1" name="PRIORITY_1"/>
				<item code="2" name="PRIORITY_2"/>
				<item code="3" name="PRIORITY_3"/>
				<item code="4" name="PRIORITY_4"/>
				<item code="5" name="PRIORITY_5"/>
				<item code="6" name="PRIORITY_6"/>
				<item code="7" name="PRIORITY_7"/>
				<item code="8" name="PRIORITY_8"/>
				<item code="9" name="PRIORITY_9"/>
				<item code="10" name="PRIORITY_10"/>
				<item code="11" name="PRIORITY_11"/>
				<item code="12" name="PRIORITY_12"/>
				<item code="13" name="PRIORITY_13"/>
				<item code="14" name="PRIORITY_14"/>
				<item code="15" name="PRIORITY_15"/>
			</data>
		</avp>

		<!-- https://www.rfc-editor.org/rfc/rfc4006.html#section-8.29 -->
		<avp name="Rating-Group" code="432" must="M" may="-" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Final-Unit-Action" code="449" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.35 -->
			<data type="Enumerated">
				<item code="0" name="TERMINATE"/>
				<item code="1" name="REDIRECT"/>
				<item code="2" name="RESTRICT_ACCESS"/>
			</data>
		</avp>

		<avp name="Redirect-Address-Type" code="433" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.38 -->
			<data type="Enumerated">
				<item code="0" name="IPv4 Address"/>
				<item code="1" name="IPv6 Address"/>
				<item code="2" name="URL"/>
				<item code="3" name="SIP URI"/>
			</data>
		</avp>

		<avp name="Redirect-Server-Address" code="435" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.39 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Subscription-Id" code="443" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.46-->
			<data type="Grouped">
				<rule avp="Subscription-Id-Type" required="true" max="1"/>
				<rule avp="Subscription-Id-Data" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Subscription-Id-Type" code="450" must="M" may="-" must-not="V" may-encrypt="Y" vendor-id="0">
            <!-- https://tools.ietf.org/rfc/rfc4006.txt -->
            <data type="Enumerated">
                <item code="0" name="END_USER_E164"/>
                <item code="1" name="END_USER_IMSI"/>
                <item code="2" name="END_USER_SIP_URI"/>
                <item code="3" name="END_USER_NAI"/>
                <item code="4" name="END_USER_PRIVATE"/>
            </data>
        </avp>

        <avp name="Subscription-Id-Data" code="444" must="M" may="-" must-not="V" may-encrypt="Y" vendor-id="0">
            <!-- https://tools.ietf.org/rfc/rfc4006.txt -->
            <data type="UTF8String"/>
        </avp>

		<avp name="User-Equipment-Info-Type" code="459" must="-" may="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.50-->
			<data type="Enumerated">
				<item code="0" name="IMEISV"/>
				<item code="1" name="MAC"/>
				<item code="2" name="EUI64"/>
				<item code="3" name="MODIFIED_EUI64"/>
			</data>
		</avp>

		<avp name="User-Equipment-Info-Value" code="460" must="-" may="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.51-->
			<data type="OctetString"/>
		</avp>

		<avp name="CC-Correlation-Id" code="411" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.1 -->
			<data type="OctetString"/>
		</avp>

		<avp name="CC-Session-Failover" code="418" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.4 -->
			<data type="Enumerated">
				<item code="0" name="FAILOVER_NOT_SUPPORTED"/>
				<item code="1" name="FAILOVER_SUPPORTED"/>
			</data>
		</avp>

		<avp name="CC-Sub-Session-Id" code="419" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.5 -->
			<data type="Unsigned64"/>
		</avp>
	
		<avp name="Check-Balance-Result" code="422" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.6 -->
			<data type="Enumerated">
				<item code="0" name="ENOUGH_CREDIT"/>
				<item code="1" name="NO_CREDIT"/>
			</data>
		</avp>

		<avp name="Cost-Information" code="423" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.7 -->
			<data type="Grouped">
				<rule avp="Unit-Value" required="true" max="1"/>
				<rule avp="Currency-Code" required="true" max="1"/>
				<rule avp="Cost-Unit" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Value-Digits" code="447" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.10-->
			<data type="Integer64"/>
		</avp>

		<avp name="Credit-Control-Failure-Handling" code="427" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.14 -->
			<data type="Enumerated">
				<item code="0" name="TERMINATE"/>
				<item code="1" name="CONTINUE"/>
				<item code="2" name="RETRY_AND_TERMINATE"/>
			</data>
		</avp>

		<avp name="Direct-Debiting-Failure-Handling" code="428" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.15 -->
			<data type="Enumerated">
				<item code="0" name="TERMINATE_OR_BUFFER"/>
				<item code="1" name="CONTINUE"/>
			</data>
		</avp>

		<avp name="Multiple-Services-Credit-Control" code="456" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Grouped">
				<rule avp="Granted-Service-Unit" required="false" max="1"/>
				<rule avp="Requested-Service-Unit" required="false" max="1"/>
				<rule avp="Used-Service-Unit" required="false"/>
				<rule avp="Service-Identifier" required="false"/>
				<rule avp="Rating-Group" required="false" max="1"/>
				<rule avp="G-S-U-Pool-Reference" required="false"/>
				<rule avp="Validity-Time" required="false" max="1"/>
				<rule avp="Result-Code" required="false" max="1"/>
				<rule avp="Final-Unit-Indication" required="false" max="1"/>
				<rule avp="Time-Quota-Threshold" required="false" max="1"/>
				<rule avp="Volume-Quota-Threshold" required="false" max="1"/>
				<rule avp="Unit-Quota-Threshold" required="false" max="1"/>
				<rule avp="Quota-Holding-Time" required="false" max="1"/>
				<rule avp="Quota-Consumption-Time" required="false" max="1"/>
				<rule avp="Reporting-Reason" required="false"/>
				<rule avp="Trigger" required="false" max="1"/>
				<rule avp="PS-Furnish-Charging-Information" required="false" max="1"/>
				<rule avp="Refund-Information" required="false" max="1"/>
				<rule avp="AF-Correlation-Information" required="false"/>
				<rule avp="Envelope" required="false"/>
				<rule avp="Envelope-Reporting" required="false" max="1"/>
				<rule avp="Time-Quota-Mechanism" required="false" max="1"/>
				<rule avp="Service-Specific-Info" required="false"/>
				<rule avp="QoS-Information" required="false" max="1"/>
				<rule avp="Announcement-Information" required="false"/>
				<rule avp="3GPP-RAT-Type" required="false" max="1"/>
				<rule avp="Related-Trigger" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Final-Unit-Indication" code="430" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.17-->
			<data type="Grouped">
				<rule avp="Final-Unit-Action" required="true" max="1"/>
				<rule avp="Restriction-Filter-Rule" required="false"/>
				<rule avp="Filter-Id" required="false"/>
				<rule avp="Redirect-Server" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="Granted-Service-Unit" code="431" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.17-->
			<data type="Grouped">
				<rule avp="Tariff-Time-Change" required="false" max="1"/>
				<rule avp="CC-Time" required="false" max="1"/>
				<rule avp="CC-Total-Octets" required="false" max="1"/>
				<rule avp="CC-Input-Octets" required="false" max="1"/>
				<rule avp="CC-Output-Octets" required="false" max="1"/>
				<rule avp="CC-Service-Specific-Units" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="Requested-Service-Unit" code="437" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.18-->
			<data type="Grouped">
				<rule avp="CC-Time" required="false" max="1"/>
				<rule avp="CC-Total-Octets" required="false" max="1"/>
				<rule avp="CC-Input-Octets" required="false" max="1"/>
				<rule avp="CC-Output-Octets" required="false" max="1"/>
				<rule avp="CC-Service-Specific-Units" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="Restriction-Filter-Rule" code="438" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.36-->
			<data type="IPFilterRule"/>
		</avp>

		<avp name="Service-Identifier" code="439" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.28-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Used-Service-Unit" code="446" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.19-->
			<data type="Grouped">
				<rule avp="Tariff-Change-Usage" required="false" max="1"/>
				<rule avp="Reporting-Reason" required="false" max="1"/>
				<rule avp="CC-Time" required="false" max="1"/>
				<rule avp="CC-Total-Octets" required="false" max="1"/>
				<rule avp="CC-Input-Octets" required="false" max="1"/>
				<rule avp="CC-Output-Octets" required="false" max="1"/>
				<rule avp="CC-Service-Specific-Units" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="G-S-U-Pool-Reference" code="457" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.30 -->
			<data type="Grouped">
				<rule avp="G-S-U-Pool-Identifier" required="true" max="1"/>
				<rule avp="CC-Unit-Type" required="true" max="1"/>
				<rule avp="Unit-Value" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Tariff-Change-Usage" code="452" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.27 -->
			<data type="Enumerated">
				<item code="0" name="UNIT_BEFORE_TARIFF_CHANGE"/>
				<item code="1" name="UNIT_AFTER_TARIFF_CHANGE"/>
				<item code="2" name="UNIT_INDETERMINATE"/>
			</data>
		</avp>

		<avp name="G-S-U-Pool-Identifier" code="453" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.31 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="CC-Unit-Type" code="454" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.32 -->
			<data type="Enumerated">
				<item code="0" name="TIME"/>
				<item code="1" name="MONEY"/>
				<item code="2" name="TOTAL-OCTETS"/>
				<item code="3" name="INPUT-OCTETS"/>
				<item code="4" name="OUTPUT-OCTETS"/>
				<item code="5" name="SERVICE-SPECIFIC-UNITS"/>
			</data>
		</avp>

		<avp name="Validity-Time" code="448" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.33-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Multiple-Services-Indicator" code="455" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.40 -->
			<data type="Enumerated">
				<item code="0" name="MULTIPLE_SERVICES_NOT_SUPPORTED"/>
				<item code="1" name="MULTIPLE_SERVICES_SUPPORTED"/>
			</data>
		</avp>

		<avp name="Requested-Action" code="436" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.41 -->
			<data type="Enumerated">
				<item code="0" name="DIRECT_DEBITING"/>
				<item code="1" name="REFUND_ACCOUNT"/>
				<item code="2" name="CHECK_BALANCE"/>
				<item code="3" name="PRICE_ENQUIRY"/>
			</data>
		</avp>

		<avp name="Cost-Unit" code="424" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.12 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Service-Context-Id" code="461" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.42-->
			<data type="UTF8String"/>
		</avp>

		<avp name="Service-Parameter-Info" code="440" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.43-->
			<data type="Grouped">
				<rule avp="Service-Parameter-Type" required="true" max="1"/>
				<rule avp="Service-Parameter-Value" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Service-Parameter-Type" code="441" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.44-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Service-Parameter-Value" code="442" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.45-->
			<data type="OctetString"/>
		</avp>

		<avp name="Tracing-Config" code="20001" must="V"	may="V" must-not="V" may-encrypt="Y" vendor-id="20858">
			<!-- Vendor Specific AVP for Casa internal tracing support -->
			<data type="Grouped"/>
				<rule avp="Trace-Ref" required="true" max="1"/>
				<rule avp="Trace-Sess-Ref" required="true" max="1"/>
				<rule avp="Trace-Span-Id" required="true" max="1"/>
				<rule avp="Trace-Id-Low" required="true" max="1"/>
				<rule avp="Trace-Id-High" required="true" max="1"/>
				<rule avp="Trace-Depth" required="true" max="1"/>
				<rule avp="Trace-UE-Id" required="true" max="1"/>
				<!-- *[ AVP ]-->
		</avp>

		<avp name="Supported-Features" code="628" vendor-id="10415" must="V,M" may-encrypt="N">
		  <data type="Grouped">
			<rule avp="Vendor-Id" required="true" max="1"/>
			<rule avp="Feature-List-ID" required="true" max="1"/>
			<rule avp="Feature-List" required="true" max="1"/>
		  </data>
		</avp>

		<avp name="AF-Charging-Identifier" code="505" must="V,M" may="P"  may-encrypt="Y" vendor-id="10415">
			    <!-- 3GPP 29.214 Section 5.3.6 -->
				<data type="OctetString"/>
		</avp>

		<avp name="Framed-IP-Address" code="8" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- https://www.rfc-editor.org/rfc/rfc4005.html#section-6.11.1 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Framed-Ipv6-Prefix" code="97" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- https://www.rfc-editor.org/rfc/rfc4005.html#section-6.11.6 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Called-Station-Id" code="30" must="M" must-not="V" may-encrypt="Y" vendor-id="0">
            <!-- https://tools.ietf.org/html/rfc4005#section-4.5-->
            <data type="UTF8String"/>
        </avp>

		<avp name="Calling-Party-Address" code="831" must="V,M" vendor-id="10415">
			    <!-- 3GPP 32.299 Section 7.2.33 -->
				<data type="UTF8String"/>
		</avp>

		<avp name="Callee-Information" code="565" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<!-- 3GPP 29.214 Section 5.3.62 -->
			<data type="Grouped">
				<rule avp="Called-Party-Address" required="true" max="1"/>
				<rule avp="Requested-Party-Address" required="false"/>
				<rule avp="Called-Asserted-Identity" required="false"/>
			</data>
		</avp>
		
		<avp name="Required-Access-Info" code="536" must="V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.214 Section 5.3.34 -->
            <data type="Enumerated">
                <item code="0" name="USER_LOCATION"/>
                <item code="1" name="MS_TIME_ZONE"/>
            </data>
        </avp>

		<avp name="Reference-Id" code="4202" must="M,V" may="P" may-encrypt="Y">
			<data type="OctetString"/>
		</avp>

        <avp name="Access-Network-Charging-Address" code="501" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP TS 29.214 Section 5.3.2 -->
            <data type="Address"/>
        </avp>

        <avp name="AN-GW-Address" code="1050" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.49 -->
            <data type="Address"/>
        </avp>

		<avp name="AN-Trusted" code="1503" must="M,V" must-not="P">
            <!-- 3GPP 29.273 5.2.3.9-->
            <data type="Enumerated">
                <item code="0" name="TRUSTED"/>
                <item code="1" name="UNTRUSTED"/>
 			</data>
        </avp>

        <avp name="IP-CAN-Type" code="1027" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.27 -->
            <data type="Enumerated">
                <item code="0" name="3GPP-GPRS"/>
                <item code="1" name="DOCSIS"/>
                <item code="2" name="xDSL"/>
                <item code="3" name="WiMAX"/>
                <item code="4" name="3GPP2"/>
                <item code="5" name="3GPP-EPS"/>
                <item code="6" name="Non-3GPP-EPS"/>
                <item code="7" name="FBA"/>
                <item code="8" name="3GPP-5GS"/>
                <item code="9" name="Non-3GPP-5GS"/>
            </data>
        </avp>

		<avp name="Flows" code="510" must="V,M"	may="P" may-encrypt="Y" vendor-id="10415">
			<!-- 3GPP 29.214 Section 5.3.10 -->
			<data type="Grouped">
				<rule avp="Media-Component-Number" required="true" max="1"/>
				<rule avp="Flow-Number" required="false"/>
				<rule avp="Content-Version" required="false"/>
				<rule avp="Final-Unit-Action" required="false" max="1"/>
				<rule avp="Media-Component-Status" required="false" max="1"/>
			</data>
		</avp>

		<avp name="NetLoc-Access-Support" code="2824" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<!-- 3GPP 29.212 5.3.111 -->
			<data type="Unsigned32"/>
		</avp>
		
		<avp name="RAT-Type" code="1032" must="V"	may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="WLAN"/>
				<item code="1" name="VIRTUAL"/>
				<item code="1000" name="UTRAN"/>
				<item code="1001" name="GERAN"/>
				<item code="1002" name="GAN"/>
				<item code="1003" name="HSPA_EVOLUTION"/>
				<item code="1004" name="EUTRAN"/>
				<item code="2000" name="CDMA2000_1X"/>
				<item code="2001" name="HRPD"/>
				<item code="2002" name="UMB"/>
				<item code="2003" name="EHRPD"/>
			</data>
		</avp>

		<avp name="User-Equipment-Info" code="458" must="-" may="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.49-->
			<data type="Grouped">
				<rule avp="User-Equipment-Info-Type" required="true" max="1"/>
				<rule avp="User-Equipment-Info-Value" required="true" max="1"/>
			</data>
		</avp>
		
        <avp name="3GPP-User-Location-Info" code="22" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.061 Section 16.4.7 -->
            <data type="OctetString"/>
        </avp>

        <avp name="3GPP-SGSN-MCC-MNC" code="18" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.061 Section 16.4.7 -->
            <data type="UTF8String"/>
        </avp>

		<avp name="User-Location-Info-Time" code="2812" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212  Section 5.3.101 -->
            <data type="Time"/>
        </avp>

		<avp name="Load" code="650" must="M,V" >
            <!-- https://tools.ietf.org/html/rfc8583#section-7.1-->
            <data type="Grouped">
                <rule avp="Load-Type" required="false" max="1"/>
                <rule avp="Load-Value" required="false" max="1"/>
                <rule avp="SourceID" required="false" max="1"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>
	
        <avp name="3GPP-MS-TimeZone" code="23" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.061 Table 9a -->
            <data type="OctetString"/>
        </avp>

		<avp name="RAN-NAS-Release-Cause" code="2819" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.106 -->
            <data type="OctetString"/>
        </avp>

		<avp name="TWAN-Identifier" code="29" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.061 Table 9a -->
            <data type="OctetString"/>
        </avp>

		<avp name="TCP-Source-Port" code="2843" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.129 -->
            <data type="Unsigned32"/>
        </avp>

		<avp name="UDP-Source-Port" code="2806" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.97 -->
            <data type="Unsigned32"/>
        </avp>

		<avp name="UE-Local-IP-Address" code="2805" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.96 -->
            <data type="Address"/>
        </avp>

		<avp name="Media-Component-Number" code="518" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.214 Section 5.3.37 -->
            <data type="Unsigned32"/>
        </avp>

		<avp name="Flow-Number" code="509" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.214 Section 5.3.9 -->
            <data type="Unsigned32"/>
        </avp>

		<avp name="Flow-Description" code="507" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
		<!-- 3GPP 29.212 -->
			<data type="IPFilterRule"/>
		</avp>

		<avp name="Max-Requested-Bandwidth-UL" code="516" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<!-- 3GPP 29.214 Section 5.3.15 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Max-Requested-Bandwidth-DL" code="515" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<!-- 3GPP 29.214 Section 5.3.14 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Extended-Max-Requested-BW-UL" code="555" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<!-- 3GPP 29.214 Section 5.3.53 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Extended-Max-Requested-BW-DL" code="554" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<!-- 3GPP 29.214 Section 5.3.52 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Flow-Status" code="511" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.214 Section 5.3.11 -->
            <data type="Enumerated">
                <item code="0" name="ENABLED-UPLINK"/>
                <item code="1" name="ENABLED-DOWNLINK"/>
                <item code="2" name="ENABLED"/>
                <item code="3" name="DISABLED"/>
                <item code="4" name="REMOVED"/>
            </data>
        </avp>

		<avp name="Max-PLR-DL" code="2852" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.138 -->
            <data type="Float32"/>
        </avp>

		<avp name="Max-PLR-UL" code="2853" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.139 -->
            <data type="Float32"/>
        </avp>
		
		<avp name="Trace-Ref" code="20002" must="V" may="V" must-not="V" may-encrypt="Y" vendor-id="20858">
			<!-- Vendor Specific AVP for Casa internal tracing support -->
			<data type="OctetString"/>
		</avp>

		<avp name="Trace-Sess-Ref" code="20003" must="V" may="V" must-not="V" may-encrypt="Y" vendor-id="20858">
			<!-- Vendor Specific AVP for Casa internal tracing support -->
			<data type="OctetString"/>
		</avp>

		<avp name="Trace-Span-Id" code="20004" must="V" may="V" must-not="V" may-encrypt="Y" vendor-id="20858">
			<!-- Vendor Specific AVP for Casa internal tracing support -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="Trace-Id-Low" code="20005" must="V" may="V" must-not="V" may-encrypt="Y" vendor-id="20858">
			<!-- Vendor Specific AVP for Casa internal tracing support -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="Trace-Id-High" code="20006" must="V" may="V" must-not="V" may-encrypt="Y" vendor-id="20858">
			<!-- Vendor Specific AVP for Casa internal tracing support -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="Trace-Depth" code="20007" must="V" may="V" must-not="V" may-encrypt="Y" vendor-id="20858">
			<!-- Vendor Specific AVP for Casa internal tracing support -->
			<data type="Enumerated">
				<item code="0" name="MINIMUM"/>
				<item code="1" name="MEDIUM"/>		
				<item code="2" name="MAXIMUM"/>		
				<item code="3" name="MINIMUM_WO_VENDOR_SPEC_EXT"/>		
				<item code="4" name="MEDIUM_WO_VENDOR_SPEC_EXT"/>		
				<item code="5" name="MAXIMUM_WO_VENDOR_SPEC_EXT"/>	
			</data>
		</avp>

		<avp name="Trace-UE-Id" code="20008" must="V" may="V" must-not="V" may-encrypt="Y" vendor-id="20858">
			<!-- Vendor Specific AVP for Casa internal tracing support -->
			<data type="OctetString"/>
		</avp>

		<avp name="Auth-Request-Type" code="274" must="M" may="P" must-not="V" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc6733#section-8.7 -->
			<data type="Enumerated">
				<item code="0" name="AUTHENTICATE_ONLY"/>
				<item code="1" name="AUTHORIZE_ONLY"/>
				<item code="2" name="AUTHORIZE_AUTHENTICATE"/>
			</data>
		</avp>

		<avp name="Service-Selection" code="493" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- https://datatracker.ietf.org/doc/html/rfc5778#section-6.2 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Idle-Timeout" code="28" must="M" may="-" must-not="V"> may-encrypt="-"
			<!-- https://datatracker.ietf.org/doc/html/rfc7155#section-4.4.4 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Framed-Pool" code="88" must="M" may="-" must-not="V" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc7155#section-4.4.10.5.4 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Framed-IPv6-Pool" code="100" must="M" may="-" must-not="V" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc7155#section-4.4.10.5.8 -->
			<data type="OctetString"/>
		</avp>

		<avp name="State" code="24" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- https://datatracker.ietf.org/doc/html/rfc4005#section-9.3.4 -->
			<data type="OctetString"/>
		</avp>

		<avp name="EAP-Payload" code="462" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc4072#section-4.1.1 -->
			<data type="OctetString"/>
		</avp>

		<avp name="MIP6-Feature-Vector" code="124" must="M" may="-" must-not="V,P" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5447#section-4.2.5 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="MIP6-Agent-Info" code="486" must="M" may="-" must-not="V,P" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5447#section-4.2.1 -->
			<data type="Grouped">
				<rule avp="MIP-Home-Agent-Address" required="false" max="2"/>
				<rule avp="MIP-Home-Agent-Host" required="false" max="1"/>
				<rule avp="MIP6-Home-Link-Prefix" required="false" max="1"/>
			</data>
		</avp>

		<avp name="MIP-Home-Agent-Address" code="334" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc4004#section-7.4 -->
			<data type="Address"/>
		</avp>

		<avp name="MIP6-Home-Link-Prefix" code="125" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc4072#section-4.2.4 -->
			<data type="OctetString"/>
		</avp>

		<avp name="MIP-Home-Agent-Host" code="348" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc4004#section-7.11 -->
			<data type="Grouped">
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Destination-Host" required="true" max="1"/>
			</data>
		</avp>

		<avp name="QoS-Capability" code="578" must="M" may="-" must-not="V,P" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.2.1 -->
			<data type="Grouped">
				<rule avp="QoS-Profile-Template" required="true" min="1"/>
			</data>
		</avp>

		<avp name="QoS-Profile-Template" code="574" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-5.3 -->
			<data type="Grouped">
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="QoS-Profile-Id" required="true" max="1"/>
			</data>
		</avp>

		<avp name="QoS-Profile-Id" code="573" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-5.2 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Visited-Network-Identifier" code="600" must="M,V" may="-" must-not="P" may-encrypt="-" vendor-id="10415">
			<!-- 29273 - 9.2.3.1.1 -->
			<data type="OctetString"/>
		</avp>

		<avp name="MIP-Careof-Address" code="487" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5778#section-6.7 -->
			<data type="Address"/>
		</avp>

		<avp name="AAA-Failure-Indication" code="1518" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- 29273 - 9.2.3.1.1 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="DER-S6b-Flags" code="1523" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- 29273 - 9.2.3.1.1 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="EAP-Master-Session-Key" code="464" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc4072#section-4.1.3 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Mobile-Node-Identifier" code="506" must="M" may="-" must-not="V,P" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5779#section-5.6 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="QoS-Resources" code="508" must="M" may="-" must-not="V,P" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-3.1 -->
			<data type="Grouped">
				<rule avp="Filter-Rule" required="true" min="1"/>
			</data>
		</avp>

		<avp name="Filter-Rule" code="509" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-3.2 -->
			<data type="Grouped">
				<rule avp="Filter-Rule-Precedence" required="false" max="1"/>
				<rule avp="Classifier" required="false" max="1"/>
				<rule avp="Time-Of-Day-Condition" required="false"/>
				<rule avp="Treatment-Action" required="false" max="1"/>
				<rule avp="QoS-Semantics" required="false" max="1"/>
				<rule avp="QoS-Profile-Template" required="false" max="1"/>
				<rule avp="QoS-Parameters" required="false" max="1"/>
				<rule avp="Excess-Treatment" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Filter-Rule-Precedence" code="510" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-3.3 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Classifier" code="511" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.1 -->
			<data type="Grouped">
				<rule avp="Classifier-ID" required="true" max="1"/>
				<rule avp="Protocol" required="false" max="1"/>
				<rule avp="Direction" required="false" max="1"/>
				<rule avp="From-Spec" required="false"/>
				<rule avp="To-Spec" required="false"/>
				<rule avp="Diffserv-Code-Point" required="false"/>
				<rule avp="IP-Option" required="false"/>
				<rule avp="TCP-Option" required="false"/>
				<rule avp="TCP-Flags" required="false" max="1"/>
				<rule avp="ICMP-Type" required="false"/>
				<rule avp="ETH-Option" required="false"/>
			</data>
		</avp>

		<avp name="Classifier-ID" code="512" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.2 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Protocol" code="513" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.3 -->
			<data type="Enumerated">
				<!-- Refer: https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml -->
				<item code="0" name="HOPOPT"/>
				<item code="1" name="ICMP"/>
				<item code="2" name="IGMP"/>
				<item code="3" name="GGP"/>
				<item code="4" name="IPv4"/>
				<item code="5" name="ST"/>
				<item code="6" name="TCP"/>
				<item code="7" name="CBT"/>
				<item code="8" name="EGP"/>
				<item code="9" name="IGP"/>
				<item code="10" name="HBBN-RCC-MON"/>
				<item code="11" name="NVP-II"/>
				<item code="12" name="PUP"/>
				<item code="13" name="ARGUS"/>
				<item code="14" name="EMCON"/>
				<item code="15" name="XNET"/>
				<item code="16" name="CHAOS"/>
				<item code="17" name="UDP"/>
				<item code="18" name="MUX"/>
				<item code="19" name="DCN-MEAS"/>
				<item code="20" name="HMP"/>
				<item code="21" name="PRM"/>
				<item code="22" name="XNS-IDP"/>
				<item code="23" name="TRUNK-1"/>
				<item code="24" name="TRUNK-2"/>
				<item code="25" name="LEAF-1"/>
				<item code="26" name="LEAF-2"/>
				<item code="27" name="RDP"/>
				<item code="28" name="IRTP"/>
				<item code="29" name="ISO-TP4"/>
				<item code="30" name="NETBLT"/>
				<item code="31" name="MFE-NSP"/>
				<item code="32" name="MERIT-INP"/>
				<item code="33" name="DCCP"/>
				<item code="34" name="3PC"/>
				<item code="35" name="IDPR"/>
				<item code="36" name="XTP"/>
				<item code="37" name="DDP"/>
				<item code="38" name="IDPR-CMTP"/>
				<item code="39" name="TP++"/>
				<item code="40" name="IL"/>
				<item code="41" name="IPv6"/>
				<item code="42" name="SDRP"/>
				<item code="43" name="IPv6-Route"/>
				<item code="44" name="IPv6-Frag"/>
				<item code="45" name="IDRP"/>
				<item code="46" name="RSVP"/>
				<item code="47" name="GRE"/>
				<item code="48" name="DSR"/>
				<item code="49" name="BNA"/>
				<item code="50" name="ESP"/>
				<item code="51" name="AH"/>
				<item code="52" name="I-NLSP"/>
				<item code="53" name="SWIPE"/>
				<item code="54" name="NARP"/>
				<item code="55" name="MOBILE"/>
				<item code="56" name="TLSP"/>
				<item code="57" name="SKIP"/>
				<item code="58" name="IPv6-ICMP"/>
				<item code="59" name="IPv6-NoNxt"/>
				<item code="60" name="IPv6-Opts"/>
				<item code="62" name="CFTP"/>
				<item code="64" name="SAT-EXPAK"/>
				<item code="65" name="KRYPTOLAN"/>
				<item code="66" name="RVD"/>
				<item code="67" name="IPPC"/>
				<item code="69" name="SAT-MON"/>
				<item code="70" name="VISA"/>
				<item code="71" name="IPCV"/>
				<item code="72" name="CPNX"/>
				<item code="73" name="CPHB"/>
				<item code="74" name="WSN"/>
				<item code="75" name="PVP"/>
				<item code="76" name="BR-SAT-MON"/>
				<item code="77" name="SUN-ND"/>
				<item code="78" name="WB-MON"/>
				<item code="79" name="WB-EXPAK"/>
				<item code="80" name="ISO-IP"/>
				<item code="81" name="VMTP"/>
				<item code="82" name="SECURE-VMTP"/>
				<item code="83" name="VINES"/>
				<item code="84" name="TTP"/>
				<item code="84" name="IPTM"/>
				<item code="85" name="NSFNET-IGP"/>
				<item code="86" name="DGP"/>
				<item code="87" name="TCF"/>
				<item code="88" name="EIGRP"/>
				<item code="89" name="OSPFIGP"/>
				<item code="90" name="Sprite-RPC"/>
				<item code="91" name="LARP"/>
				<item code="92" name="MTP"/>
				<item code="93" name="AX.25"/>
				<item code="94" name="IPIP"/>
				<item code="96" name="SCC-SP"/>
				<item code="97" name="ETHERIP"/>
				<item code="98" name="ENCAP"/>
				<item code="100" name="GMTP"/>
				<item code="101" name="IFMP"/>
				<item code="102" name="PNNI"/>
				<item code="103" name="PIM"/>
				<item code="104" name="ARIS"/>
				<item code="105" name="SCPS"/>
				<item code="106" name="QNX"/>
				<item code="107" name="A/N"/>
				<item code="108" name="IPComp"/>
				<item code="109" name="SNP"/>
				<item code="110" name="Compaq-Peer"/>
				<item code="111" name="IPX-in-IP"/>
				<item code="112" name="VRRP"/>
				<item code="113" name="PGM"/>
				<item code="115" name="L2TP"/>
				<item code="116" name="DDX"/>
				<item code="117" name="IATP"/>
				<item code="118" name="STP"/>
				<item code="119" name="SRP"/>
				<item code="120" name="UTI"/>
				<item code="121" name="SMP"/>
				<item code="123" name="PTP"/>
				<item code="124" name="ISIS"/>
				<item code="125" name="FIRE"/>
				<item code="126" name="CRTP"/>
				<item code="127" name="CRUDP"/>
				<item code="128" name="SSCOPMCE"/>
				<item code="129" name="IPLT"/>
				<item code="130" name="SPS"/>
				<item code="131" name="PIPE"/>
				<item code="132" name="SCTP"/>
				<item code="133" name="FC"/>
				<item code="134" name="RSVP-E2E-IGNORE"/>
				<item code="135" name="Mobility"/>
				<item code="136" name="UDPLite"/>
				<item code="137" name="MPLS-in-IP"/>
				<item code="138" name="manet"/>
				<item code="139" name="HIP"/>
				<item code="140" name="Shim6"/>
				<item code="141" name="WESP"/>
				<item code="142" name="ROHC"/>
				<item code="143" name="Ethernet"/>
				<item code="144" name="AGGFRAG"/>
				<item code="255" name="Reserved"/>
			</data>
		</avp>

		<avp name="Direction" code="514" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.4 -->
			<data type="Enumerated">
				<item code="0" name="IN"/>
				<item code="1" name="OUT"/>
				<item code="2" name="BOTH"/>
			</data>
		</avp>

		<avp name="From-Spec" code="515" must="-" may="-" must-not="V" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.5 -->
			<data type="Grouped">
				<rule avp="IP-Address" required="false"/>
				<rule avp="IP-Address-Range" required="false"/>
				<rule avp="IP-Address-Mask" required="false"/>
				<rule avp="MAC-Address" required="false"/>
				<rule avp="MAC-Address-Mask" required="false"/>
				<rule avp="EUI64-Address" required="false"/>
				<rule avp="EUI64-Address-Mask" required="false"/>
				<rule avp="Port" required="false"/>
				<rule avp="Port-Range" required="false"/>
				<rule avp="Negated" required="false" max="1"/>
				<rule avp="Use-Assigned-Address" required="false" max="1"/>
			</data>
		</avp>	

		<avp name="To-Spec" code="516" must="-" may="-" must-not="V" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.6 -->
			<data type="Grouped">
				<rule avp="IP-Address" required="false"/>
				<rule avp="IP-Address-Range" required="false"/>
				<rule avp="IP-Address-Mask" required="false"/>
				<rule avp="MAC-Address" required="false"/>
				<rule avp="MAC-Address-Mask" required="false"/>
				<rule avp="EUI64-Address" required="false"/>
				<rule avp="EUI64-Address-Mask" required="false"/>
				<rule avp="Port" required="false"/>
				<rule avp="Port-Range" required="false"/>
				<rule avp="Negated" required="false" max="1"/>
				<rule avp="Use-Assigned-Address" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Diffserv-Code-Point" code="535" must="-" may="-" must-not="-" may-encrypt="-" vendor-id="0">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.1 -->
			<data type="Enumerated">
				<!-- mazhenfeng: sorry, I can not found define in spec and rfc, please improve here if you can found -->
			</data>
		</avp>
		
		<avp name="IP-Address" code="518" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.2 -->
			<data type="Address"/>
		</avp>

		<avp name="IP-Address-Range" code="519" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.3 -->
			<data type="Grouped">
				<rule avp="IP-Address-Start" required="false" max="1"/>
				<rule avp="IP-Address-End" required="false" max="1"/>
			</data>
		</avp>

		<avp name="IP-Address-Start" code="520" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.4 -->
			<data type="Address"/>
		</avp>

		<avp name="IP-Address-End" code="521" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.5 -->
			<data type="Address"/>
		</avp>

		<avp name="IP-Address-Mask" code="522" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.6 -->
			<data type="Grouped">
				<rule avp="IP-Address" required="false" max="1"/>
				<rule avp="IP-Bit-Mask-Width" required="false" max="1"/>
			</data>
		</avp>

		<avp name="IP-Bit-Mask-Width" code="523" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.7 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="MAC-Address" code="524" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.8 -->
			<data type="OctetString"/>
		</avp>

		<avp name="MAC-Address-Mask" code="525" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.9 -->
			<data type="Grouped">
				<rule avp="MAC-Address" required="false" max="1"/>
				<rule avp="MAC-Address-Mask-Pattern" required="false" max="1"/>
			</data>
		</avp>

		<avp name="MAC-Address-Mask-Pattern" code="526" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.10 -->
			<data type="OctetString"/>
		</avp>

		<avp name="EUI64-Address" code="527" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.11 -->
			<data type="OctetString"/>
		</avp>

		<avp name="EUI64-Address-Mask" code="528" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.12 -->
			<data type="Grouped">
				<rule avp="EUI64-Address" required="false" max="1"/>
				<rule avp="EUI64-Address-Mask-Pattern" required="false" max="1"/>
			</data>
		</avp>

		<avp name="EUI64-Address-Mask-Pattern" code="529" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.13 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Port" code="530" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.14 -->
			<data type="Integer32"/>
		</avp>

		<avp name="Port-Range" code="531" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.15 -->
			<data type="Grouped">
				<rule avp="Port-Start" required="false" max="1"/>
				<rule avp="Port-End" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Port-Start" code="532" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.15 -->
			<data type="Integer32"/>
		</avp>

		<avp name="Port-End" code="533" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.16 -->
			<data type="Integer32"/>
		</avp>

		<avp name="Negated" code="517" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.1 -->
			<data type="Enumerated">
				<item code="0" name="False"/>
				<item code="1" name="True"/>
			</data>
		</avp>

		<avp name="Use-Assigned-Address" code="534" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.7.18 -->
			<data type="Enumerated">
				<item code="0" name="False"/>
				<item code="1" name="True"/>
			</data>
		</avp>

		<avp name="IP-Option" code="537" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.3 -->
			<data type="Grouped">
				<rule avp="IP-Option-Type" required="true" max="1"/>
				<rule avp="IP-Option-Value" required="false"/>
				<rule avp="Negated" required="false" max="1"/>
			</data>
		</avp>

		<avp name="IP-Option-Type" code="538" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.4 -->
			<data type="Enumerated">
				<!-- Refer: https://www.iana.org/assignments/ip-parameters/ip-parameters.xhtml -->
				<item code="0" name="End of Option List"/>
				<item code="1" name="No-Operation"/>
				<item code="7" name="Record Route"/>
				<item code="10" name="Experimental Measurement"/>
				<item code="11" name="MTU Probe"/>
				<item code="12" name="MTU Reply"/>
				<item code="25" name="Quick-Start"/>
				<item code="82" name="Traceroute"/>
				<item code="130" name="Security"/>
				<item code="131" name="Loose Source Route"/>
				<item code="133" name="Extended Security"/>
				<item code="134" name="Commercial Security"/>
				<item code="136" name="Stream ID"/>
				<item code="137" name="Strict Source Route"/>
				<item code="142" name="Experimental Access Control"/>
				<item code="144" name="IMI Traffic Descriptor"/>
				<item code="145" name="Extended Internet Protocol"/>
				<item code="147" name="Address Extension"/>
				<item code="148" name="Router Alert"/>
				<item code="149" name="Selective Directed Broadcast"/>
				<item code="151" name="Dynamic Packet State"/>
				<item code="152" name="Upstream Multicast Pkt"/>
				<item code="205" name="Experimental Flow Control"/>
			</data>
		</avp>

		<avp name="IP-Option-Value" code="539" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.5 -->
			<data type="OctetString"/>
		</avp>

		<avp name="TCP-Option" code="540" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.6 -->
			<data type="Grouped">
				<rule avp="TCP-Option-Type" required="true" max="1"/>
				<rule avp="TCP-Option-Value" required="false"/>
				<rule avp="Negated" required="false" max="1"/>
			</data>
		</avp>

		<avp name="TCP-Option-Type" code="541" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.7 -->
			<data type="Enumerated">
				<!-- Refer: https://www.iana.org/assignments/tcp-parameters/tcp-parameters.xhtml#tcp-parameters-1 -->
				<item code="0" name="End of Option List"/>
				<item code="1" name="No-Operation"/>
				<item code="2" name="Maximum Segment Size"/>
				<item code="3" name="Window Scale"/>
				<item code="4" name="SACK Permitted"/>
				<item code="5" name="SACK"/>
				<item code="6" name="Echo"/>
				<item code="7" name="Echo Reply"/>
				<item code="8" name="Timestamps"/>
				<item code="16" name="Skeeter"/>
				<item code="17" name="Bubba"/>
				<item code="18" name="Trailer Checksum Option"/>
				<item code="19" name="MD5 Signature Option"/>
				<item code="20" name="SCPS Capabilities"/>
				<item code="21" name="Selective Negative Acknowledgements"/>
				<item code="22" name="Record Boundaries"/>
				<item code="23" name="Corruption experienced"/>
				<item code="24" name="SNAP"/>
				<item code="26" name="TCP Compression Filter"/>
				<item code="27" name="Quick-Start Response"/>
				<item code="28" name="User Timeout Option"/>
				<item code="29" name="TCP Authentication Option"/>
				<item code="30" name="Multipath TCP"/>
				<item code="34" name="TCP Fast Open Cookie"/>
				<item code="69" name="Encryption Negotiation"/>
				<item code="172" name="Accurate ECN Order 0"/>
				<item code="174" name="Accurate ECN Order 1"/>
				<item code="253" name="RFC3692-style Experiment 1"/>
				<item code="254" name="RFC3692-style Experiment 2"/>
			</data>
		</avp>

		<avp name="TCP-Option-Value" code="542" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.8 -->
			<data type="OctetString"/>
		</avp>

		<avp name="TCP-Flags" code="543" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.9 -->
			<data type="Grouped">
				<rule avp="TCP-Flag-Type" required="true" max="1"/>
				<rule avp="Negated" required="false" max="1"/>
			</data>
		</avp>

		<avp name="TCP-Flag-Type" code="544" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.10 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="ICMP-Type" code="545" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.11 -->
			<data type="Grouped">
				<rule avp="ICMP-Type-Number" required="true" max="1"/>
				<rule avp="ICMP-Code" required="false"/>
				<rule avp="Negated" required="false" max="1"/>
			</data>
		</avp>

		<avp name="ICMP-Type-Number" code="546" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.12 -->
			<data type="Enumerated">
				<!-- Refer: https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml#icmp-parameters-types -->
				<item code="0" name="Echo Reply"/>
				<item code="3" name="Destination Unreachable"/>
				<item code="5" name="Redirect"/>
				<item code="7" name="Unassigned"/>
				<item code="8" name="Echo"/>
				<item code="9" name="Router Advertisement"/>
				<item code="10" name="Router Solicitation"/>
				<item code="11" name="Time Exceeded"/>
				<item code="12" name="Parameter Problem"/>
				<item code="13" name="Timestamp"/>
				<item code="14" name="Timestamp Reply"/>
				<item code="40" name="Photuris"/>
				<item code="41" name="ICMP messages utilized by experimental mobility protocols such as Seamoby"/>
				<item code="42" name="Extended Echo Request"/>
				<item code="43" name="Extended Echo Reply"/>
				<item code="253" name="RFC3692-style Experiment 1"/>
				<item code="254" name="RFC3692-style Experiment 2"/>
				<item code="255" name="Reserved"/>
			</data>
		</avp>

		<avp name="ICMP-Code" code="547" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.13 -->
			<data type="Enumerated">
				<!-- Can not search define from RFC -->
				<!-- Refer: https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml#icmp-parameters-types -->
				<!-- meaning of ICMP-Code is depended on ICMP-Type-Number -->
				<item code="0" name="0"/>
				<item code="1" name="1"/>
				<item code="2" name="2"/>
				<item code="3" name="3"/>
				<item code="4" name="4"/>
				<item code="5" name="5"/>
				<item code="6" name="6"/>
				<item code="7" name="7"/>
				<item code="8" name="8"/>
				<item code="9" name="9"/>
				<item code="10" name="10"/>
				<item code="11" name="11"/>
				<item code="12" name="12"/>
				<item code="13" name="13"/>
				<item code="14" name="14"/>
				<item code="15" name="15"/>
			</data>
		</avp>

		<avp name="ETH-Option" code="548" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.14 -->
			<data type="Grouped">
				<rule avp="ETH-Proto-Type" required="true" max="1"/>
				<rule avp="VLAN-ID-Range" required="false"/>
				<rule avp="User-Priority-Range" required="false"/>
			</data>
		</avp>

		<avp name="ETH-Proto-Type" code="549" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.15 -->
			<data type="Grouped">
				<rule avp="ETH-Ether-Type" required="false"/>
				<rule avp="ETH-SAP" required="false"/>
			</data>
		</avp>

		<avp name="ETH-Ether-Type" code="550" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.16 -->
			<data type="OctetString"/>
		</avp>

		<avp name="ETH-SAP" code="551" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.17 -->
			<data type="OctetString"/>
		</avp>

		<avp name="VLAN-ID-Range" code="552" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.18 -->
			<data type="Grouped">
				<rule avp="S-VID-Start" required="false" max="1"/>
				<rule avp="S-VID-End" required="false" max="1"/>
				<rule avp="C-VID-Start" required="false" max="1"/>
				<rule avp="C-VID-End" required="false" max="1"/>
			</data>
		</avp>

		<avp name="S-VID-Start" code="553" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.19 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="S-VID-End" code="554" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.20 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="C-VID-Start" code="555" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.21 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="C-VID-End" code="556" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.22 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="User-Priority-Range" code="557" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.23 -->
			<data type="Grouped">
				<rule avp="Low-User-Priority" required="false"/>
				<rule avp="High-User-Priority" required="false"/>
			</data>
		</avp>

		<avp name="Low-User-Priority" code="558" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.24 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="High-User-Priority" code="559" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.1.8.25 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Time-Of-Day-Condition" code="560" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.2.1 -->
			<data type="Grouped">
				<rule avp="Time-Of-Day-Start" required="false" max="1"/>
				<rule avp="Time-Of-Day-End" required="false" max="1"/>
				<rule avp="Day-Of-Week-Mask" required="false" max="1"/>
				<rule avp="Day-Of-Month-Mask" required="false" max="1"/>
				<rule avp="Month-Of-Year-Mask" required="false" max="1"/>
				<rule avp="Absolute-Start-Time" required="false" max="1"/>
				<rule avp="Absolute-End-Time" required="false" max="1"/>
				<rule avp="Timezone-Flag" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Time-Of-Day-Start" code="561" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.2.2 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Time-Of-Day-End" code="562" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.2.3 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Day-Of-Week-Mask" code="563" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.2.4 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Day-Of-Month-Mask" code="564" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.2.5 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Month-Of-Year-Mask" code="565" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.2.6 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Absolute-Start-Time" code="566" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.2.7 -->
			<data type="Time"/>
		</avp>

		<avp name="Absolute-End-Time" code="568" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.2.9 -->
			<data type="Time"/>
		</avp>

		<avp name="Timezone-Flag" code="570" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-4.2.11 -->
			<data type="Enumerated">
				<item code="0" name="UTC"/>
				<item code="1" name="LOCAL"/>
				<item code="2" name="OFFSET"/>
			</data>
		</avp>

		<avp name="Treatment-Action" code="572" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-5.1 -->
			<data type="Enumerated">
				<item code="0" name="DROP"/>
				<item code="1" name="SHAPE"/>
				<item code="2" name="MARK"/>
				<item code="3" name="PERMIT"/>
			</data>
		</avp>

		<avp name="QoS-Profile-Template" code="574" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-5.3 -->
			<data type="Grouped">
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="QoS-Profile-Id" required="false" max="1"/>
			</data>
		</avp>

		<avp name="QoS-Semantics" code="575" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-5.4 -->
			<data type="Enumerated">
				<item code="0" name="QoS-Desired"/>
				<item code="1" name="QoS-Available"/>
				<item code="2" name="QoS-Delivered"/>
				<item code="3" name="Minimum-QoS"/>
				<item code="4" name="QoS-Authorized"/>
			</data>
		</avp>

		<avp name="QoS-Parameters" code="576" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-5.5 -->
			<data type="Grouped">
				<!-- https://www.rfc-editor.org/rfc/rfc5624 -->
				<rule avp="TMOD-1" required="false" max="1"/>
				<rule avp="TMOD-2" required="false" max="1"/>
				<rule avp="Bandwidth" required="false" max="1"/>
				<rule avp="PHB-Class" required="false" max="1"/>
			</data>
		</avp>

		<avp name="TMOD-1" code="495" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://www.rfc-editor.org/rfc/rfc5624#section-3.1 -->
			<data type="Grouped">
				<!-- https://www.rfc-editor.org/rfc/rfc5624 -->
				<rule avp="Token-Rate" required="true" max="1"/>
				<rule avp="Bucket-Depth" required="true" max="1"/>
				<rule avp="Peak-Traffic-Rate" required="true" max="1"/>
				<rule avp="Minimum-Policed-Unit" required="true" max="1"/>
				<rule avp="Maximum-Packet-Size" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Token-Rate" code="496" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://www.rfc-editor.org/rfc/rfc5624#section-3.1.1 -->
			<data type="Float32"/>
		</avp>

		<avp name="Bucket-Depth" code="497" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://www.rfc-editor.org/rfc/rfc5624#section-3.1.2 -->
			<data type="Float32"/>
		</avp>

		<avp name="Peak-Traffic-Rate" code="498" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://www.rfc-editor.org/rfc/rfc5624#section-3.1.3 -->
			<data type="Float32"/>
		</avp>

		<avp name="Minimum-Policed-Unit" code="499" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://www.rfc-editor.org/rfc/rfc5624#section-3.1.4 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Maximum-Packet-Size" code="500" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://www.rfc-editor.org/rfc/rfc5624#section-3.1.5 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="TMOD-2" code="501" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://www.rfc-editor.org/rfc/rfc5624#section-3.2 -->
			<data type="Grouped">
				<!-- https://www.rfc-editor.org/rfc/rfc5624 -->
				<rule avp="Token-Rate" required="true" max="1"/>
				<rule avp="Bucket-Depth" required="true" max="1"/>
				<rule avp="Peak-Traffic-Rate" required="true" max="1"/>
				<rule avp="Minimum-Policed-Unit" required="true" max="1"/>
				<rule avp="Maximum-Packet-Size" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Bandwidth" code="502" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://www.rfc-editor.org/rfc/rfc5624#section-3.3 -->
			<data type="Float32"/>
		</avp>

		<avp name="PHB-Class" code="503" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://www.rfc-editor.org/rfc/rfc5624#section-3.4 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Excess-Treatment" code="577" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc5777#section-5.6 -->
			<data type="Grouped">
				<rule avp="Treatment-Action" required="true" max="1"/>
				<rule avp="QoS-Profile-Template" required="false" max="1"/>
				<rule avp="QoS-Parameters" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Origination-Time-Stamp" code="1536" must="V" may="-" must-not="M,P" may-encrypt="-" vendor-id="10415">
			<!-- 29273 - 9.2.3.2.1 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="Maximum-Wait-Time" code="1537" must="V" may="-" must-not="M,P" may-encrypt="-" vendor-id="10415">
			<!-- 29273 - 9.2.3.2.1 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Emergency-Services" code="1538" must="V" may="-" must-not="M,P" may-encrypt="-" vendor-id="10415">
			<!-- 29273 - 9.2.3.2.1 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="RAR-Flags" code="1522" must="V" may="-" must-not="M,P" may-encrypt="-">
			<!-- 29273 - 9.2.3.2.1 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Trace-Info" code="1505" must="V" may="-" must-not="M,P" may-encrypt="-">
			<!-- 29273 - 8.2.3.13 -->
			<data type="Grouped">
				<rule avp="Trace-Data" required="false" max="1"/>
				<rule avp="Trace-Reference" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Trace-Data" code="1458" must="M,V" may-encrypt="N" vendor-id="10415">
			<!-- 3GPP TS 29.272 Section 7.3.63 -->
			<data type="Grouped">
				<rule avp="Trace-Reference" required="true" max="1"/>
				<rule avp="Trace-Depth" required="true" max="1"/>
				<rule avp="Trace-NE-Type-List" required="true" max="1"/>
				<rule avp="Trace-Interface-List" required="false" max="1"/>
				<rule avp="Trace-Event-List" required="true" max="1"/>
				<rule avp="OMC-Id" required="false" max="1"/>
				<rule avp="Trace-Collection-Entity" required="true" max="1"/>
				<rule avp="MDT-Configuration" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="Trace-Reference" code="1459" must="M,V" may-encrypt="N" vendor-id="10415">
			<!-- 29272 - 7.3.64 -->
            <data type="OctetString"/>
        </avp>

		<avp name="Trace-NE-Type-List" code="1463" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.68 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Interface-List" code="1464" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.69 -->
            <data type="OctetString"/>
        </avp>

		<avp name="Trace-Event-List" code="1465" must="M,V" may="-" must-not="-" may-encrypt="N">
			<!-- 29272 - 7.3.70 -->
			<data type="OctetString"/>
		</avp>

		<avp name="OMC-Id" code="1466" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Collection-Entity" code="1452" must="M,V" may-encrypt="N" vendor-id="10415">
            <data type="Address"/>
        </avp>

		<avp name="MDT-Configuration" code="1622" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.136 -->
            <data type="Grouped">
                <rule avp="Job-Type" required="true" max="1"/>
                <rule avp="Area-Scope" required="false" max="1"/>
                <rule avp="List-Of-Measurements" required="false" max="1"/>
                <rule avp="Reporting-Trigger" required="false" max="1"/>
                <rule avp="Report-Interval" required="false" max="1"/>
                <rule avp="Report-Amount" required="false" max="1"/>
                <rule avp="Event-Threshold-RSRP" required="false" max="1"/>
                <rule avp="Event-Threshold-RSRQ" required="false" max="1"/>
                <rule avp="Logging-Interval" required="false" max="1"/>
                <rule avp="Logging-Duration" required="false" max="1"/>
                <rule avp="Measurement-Period-LTE" required="false" max="1"/>
                <rule avp="Measurement-Period-UMTS" required="false" max="1"/>
                <rule avp="Collection-Period-RRM-LTE" required="false" max="1"/>
                <rule avp="Collection-Period-RRM-UMTS" required="false" max="1"/>
                <rule avp="Positioning-Method" required="false" max="1"/>
                <rule avp="Measurement-Quantity" required="false" max="1"/>
                <rule avp="Event-Threshold-Event-1F" required="false" max="1"/>
                <rule avp="Event-Threshold-Event-1I" required="false" max="1"/>
                <rule avp="MDT-Allowed-PLMN-Id" required="false"/>
                <rule avp="MBSFN-Area" required="false"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

		<avp name="MBSFN-Area" code="1694" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Grouped">
                <rule avp="MBSFN-Area-ID" required="false" max="1"/>
                <rule avp="Carrier-Frequency" required="false" max="1"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

		<avp name="MBSFN-Area-ID" code="1695" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

		<avp name="Carrier-Frequency" code="1696" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

		<avp name="Job-Type" code="1623" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="Immediate MDT only"/>
                <item code="1" name="Logged MDT only"/>
                <item code="2" name="Trace only"/>
                <item code="3" name="Immediate MDT and Trace"/>
                <item code="4" name="RLF reports only"/>
                <item code="5" name="RCEF reports only"/>
                <item code="6" name="Logged MBSFN MDT"/>
            </data>
        </avp>

		<avp name="Area-Scope" code="1624" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
			<!-- 29272 - 7.3.138 -->
			<data type="Grouped">
                <rule avp="Cell-Global-Identity" required="false"/>
                <rule avp="E-UTRAN-Cell-Global-Identity" required="false"/>
                <rule avp="Routing-Area-Identity" required="false"/>
				<rule avp="Location-Area-Identity" required="false"/>
				<rule avp="Tracking-Area-Identity" required="false"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

		<avp name="List-Of-Measurements" code="1625" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

		<avp name="Reporting-Trigger" code="1626" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

		<avp name="Report-Interval" code="1627" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
			<!-- 29272 - 7.3.141 -->
			<data type="Enumerated">
                <item code="0" name="UTMTS_250ms"/>
                <item code="1" name="UTMTS_500ms"/>
                <item code="2" name="UTMTS_1000ms"/>
                <item code="3" name="UTMTS_2000ms"/>
                <item code="4" name="UTMTS_3000ms"/>
                <item code="5" name="UTMTS_4000ms"/>
                <item code="6" name="UTMTS_6000ms"/>
                <item code="7" name="UTMTS_8000ms"/>
                <item code="8" name="UTMTS_12000ms"/>
                <item code="9" name="UTMTS_16000ms"/>
                <item code="10" name="UTMTS_20000ms"/>
                <item code="11" name="UTMTS_24000ms"/>
                <item code="12" name="UTMTS_28000ms"/>
                <item code="13" name="UTMTS_32000ms"/>
                <item code="14" name="UTMTS_64000ms"/>
                <item code="15" name="LTE_120ms"/>
                <item code="16" name="LTE_240ms"/>
                <item code="17" name="LTE_480ms"/>
                <item code="18" name="LTE_640ms"/>
                <item code="19" name="LTE_1024ms"/>
                <item code="20" name="LTE_2048ms"/>
                <item code="21" name="LTE_5120ms"/>
                <item code="22" name="LTE_10240ms"/>
                <item code="23" name="LTE_60000ms"/>
                <item code="24" name="LTE_360000ms"/>
                <item code="25" name="LTE_720000ms"/>
                <item code="26" name="LTE_1800000ms"/>
                <item code="27" name="LTE_3600000ms"/>
                <item code="28" name="NR_120ms"/>
                <item code="29" name="NR_240ms"/>
                <item code="30" name="NR_480ms"/>
                <item code="31" name="NR_640ms"/>
                <item code="32" name="NR_1024ms"/>
                <item code="33" name="NR_2048ms"/>
                <item code="34" name="NR_5120ms"/>
                <item code="35" name="NR_10240ms"/>
                <item code="36" name="NR_20480ms"/>
                <item code="37" name="NR_40960ms"/>
                <item code="38" name="NR_60000ms"/>
                <item code="39" name="NR_360000ms"/>
                <item code="40" name="NR_720000ms"/>
                <item code="41" name="NR_1800000ms"/>
                <item code="42" name="NR_3600000ms"/>
            </data>
        </avp>

		<avp name="Report-Amount" code="1628" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
			<!-- 29272 - 7.3.142 -->
			<data type="Enumerated">
                <item code="0" name="1"/>
                <item code="1" name="2"/>
                <item code="2" name="4"/>
                <item code="3" name="8"/>
                <item code="4" name="16"/>
                <item code="5" name="32"/>
                <item code="6" name="64"/>
                <item code="7" name="infinity"/>
            </data>
        </avp>

		<avp name="Event-Threshold-RSRP" code="1629" must="V" may="-" must-not="M" may-encrypt="N">
			<!-- 29272 - 7.3.143 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Event-Threshold-RSRQ" code="1630" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

		<avp name="Logging-Interval" code="1631" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="UMTS_OR_LTE_OR_NR_1280ms "/>
                <item code="1" name="UMTS_OR_LTE_OR_NR_2560ms "/>
                <item code="2" name="UMTS_OR_LTE_OR_NR_5120ms "/>
                <item code="3" name="UMTS_OR_LTE_OR_NR_10240ms"/>
                <item code="4" name="UMTS_OR_LTE_OR_NR_20480ms"/>
                <item code="5" name="UMTS_OR_LTE_OR_NR_30720ms"/>
                <item code="6" name="UMTS_OR_LTE_OR_NR_40960ms"/>
                <item code="7" name="UMTS_OR_LTE_OR_NR_61440ms"/>
                <item code="8" name="ONLY_NR_320ms"/>
                <item code="9" name="ONLY_NR_640ms"/>
                <item code="10" name="ONLY_NR_Infinity"/>
            </data>
        </avp>

		<avp name="Logging-Duration" code="1632" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="UTMS_OR_LTE_600s"/>
                <item code="1" name="UTMS_OR_LTE_1200s"/>
                <item code="2" name="UTMS_OR_LTE_2400s"/>
                <item code="3" name="UTMS_OR_LTE_3600s"/>
                <item code="4" name="UTMS_OR_LTE_5400s"/>
                <item code="5" name="UTMS_OR_LTE_7200s"/>
                <item code="6" name="NR_600s"/>
                <item code="7" name="NR_1200s"/>
                <item code="8" name="NR_2400s"/>
                <item code="9" name="NR_3600s"/>
                <item code="10" name="NR_5400s"/>
                <item code="11" name="NR_7200s"/>
            </data>
        </avp>

		<avp name="Measurement-Period-LTE" code="1655" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="1024ms"/>
                <item code="1" name="1280ms"/>
                <item code="2" name="2048ms"/>
                <item code="3" name="2560ms"/>
                <item code="4" name="5120ms"/>
                <item code="5" name="10240ms"/>
                <item code="6" name="60000ms"/>
            </data>
        </avp>

		<avp name="Measurement-Period-UMTS" code="1656" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="250ms"/>
                <item code="1" name="500ms"/>
                <item code="2" name="1000ms"/>
                <item code="3" name="2000ms"/>
                <item code="4" name="3000ms"/>
                <item code="5" name="4000ms"/>
                <item code="6" name="6000ms"/>
                <item code="7" name="8000ms"/>
                <item code="8" name="12000ms"/>
                <item code="9" name="16000ms"/>
                <item code="10" name="20000ms"/>
                <item code="11" name="24000ms"/>
                <item code="12" name="28000ms"/>
                <item code="13" name="32000ms"/>
                <item code="14" name="64000ms"/>
            </data>
        </avp>

		<avp name="Collection-Period-RRM-LTE" code="1657" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="1024ms"/>
                <item code="1" name="1280ms"/>
                <item code="2" name="2048ms"/>
                <item code="3" name="2560ms"/>
                <item code="4" name="5120ms"/>
                <item code="5" name="10240ms"/>
                <item code="6" name="60000ms"/>
            </data>
        </avp>

		<avp name="Collection-Period-RRM-UMTS" code="1658" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Enumerated">
                <item code="0" name="250ms"/>
                <item code="1" name="500ms"/>
                <item code="2" name="1000ms"/>
                <item code="3" name="2000ms"/>
                <item code="4" name="3000ms"/>
                <item code="5" name="4000ms"/>
                <item code="6" name="6000ms"/>
                <item code="7" name="8000ms"/>
                <item code="8" name="12000ms"/>
                <item code="9" name="16000ms"/>
                <item code="10" name="20000ms"/>
                <item code="11" name="24000ms"/>
                <item code="12" name="28000ms"/>
                <item code="13" name="32000ms"/>
                <item code="14" name="64000ms"/>
            </data>
        </avp>

		<avp name="Positioning-Method" code="1659" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

		<avp name="Measurement-Quantity" code="1660" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

		<avp name="Event-Threshold-Event-1F" code="1661" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Integer32"/>
        </avp>

		<avp name="Event-Threshold-Event-1I" code="1662" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="Integer32"/>
        </avp>

		<avp name="MDT-Allowed-PLMN-Id" code="1671" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <data type="OctetString"/>
        </avp>

		<avp name="RAR-Flags" code="1522" must="V" may="-" must-not="M,P" may-encrypt="-">
			<!-- 29273 - 9.2.3.1.5 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="APN-Configuration" code="1430" must="M,V" may="-" must-not="P" may-encrypt="N">
			<!-- 29272 7.3.35 -->
			<data type="Grouped">
				<rule avp="Context-Identifier" required="true" max="1"/>
				<rule avp="Served-Party-IP-Address" required="false" max="2"/>
				<rule avp="PDN-Type" required="true" max="1"/>
				<rule avp="Service-Selection" required="true" max="1"/>
				<rule avp="EPS-Subscribed-QoS-Profile" required="false" max="1"/>
				<rule avp="VPLMN-Dynamic-Address-Allowed" required="false" max="1"/>
				<rule avp="MIP6-Agent-Info" required="false" max="1"/>
				<rule avp="Visited-Network-Identifier" required="false" max="1"/>
				<rule avp="PDN-GW-Allocation-Type" required="false" max="1"/>
				<rule avp="3GPP-Charging-Characteristics" required="false" max="1"/>
				<rule avp="AMBR" required="false" max="1"/>
				<rule avp="Specific-APN-Info" required="false"/>
				<rule avp="APN-OI-Replacement" required="false" max="1"/>
				<rule avp="SIPTO-Permission" required="false" max="1"/>
				<rule avp="LIPA-Permission" required="false" max="1"/>
				<rule avp="Restoration-Priority" required="false" max="1"/>
				<rule avp="SIPTO-Local-Network-Permission" required="false" max="1"/>
				<rule avp="WLAN-offloadability" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Context-Identifier" code="1423" must="M,V" may="-" must-not="-" may-encrypt="N">
			<!-- 29272 7.3.27 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="3GPP-Charging-Characteristics" code="13" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<!-- 29061 - 16a.5 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Served-Party-IP-Address" code="848" must="V,M" may="-" must-not="-" may-encrypt="-">
			<!-- 32299 7.2.0 -->
			<data type="Address"/>
		</avp>

		<avp name="PDN-Type" code="1456" must="M,V" may="-" must-not="-" may-encrypt="N">
			<!-- 29272 - 7.3.62 -->
			<data type="Enumerated">
				<item code="0" name="IPv4"/>
				<item code="1" name="IPv6"/>
				<item code="2" name="IPv4v6"/>
				<item code="3" name="IPv4_OR_IPv6"/>
			</data>
		</avp>

		<avp name="Service-Selection" code="493" must="M,V" may="-" must-not="-" may-encrypt="-">
			<!-- 29272 - 7.3.36 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="EPS-Subscribed-QoS-Profile" code="1431" must="M,V" may="-" must-not="-" may-encrypt="N">
			<!-- 29272 7.3.37 -->
			<data type="Grouped">
				<rule avp="QoS-Class-Identifier" required="true" max="1"/>
				<rule avp="Allocation-Retention-Priority" required="true" max="1"/>
				<rule avp="PDN-Type" required="true" max="1"/>
			</data>
		</avp>

		<avp name="QoS-Class-Identifier" code="1028" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="Reserved"/>
				<item code="1" name="QCI_1"/>
				<item code="2" name="QCI_2"/>
				<item code="3" name="QCI_3"/>
				<item code="4" name="QCI_4"/>
				<item code="5" name="QCI_5"/>
				<item code="6" name="QCI_6"/>
				<item code="7" name="QCI_7"/>
				<item code="8" name="QCI_8"/>
				<item code="9" name="QCI_9"/>
				<item code="65" name="QCI_65"/>
				<item code="66" name="QCI_66"/>
				<item code="67" name="QCI_67"/>
				<item code="69" name="QCI_69"/>
				<item code="70" name="QCI_70"/>
				<item code="71" name="QCI_71"/>
				<item code="72" name="QCI_72"/>
				<item code="73" name="QCI_73"/>
				<item code="74" name="QCI_74"/>
				<item code="75" name="QCI_75"/>
				<item code="76" name="QCI_76"/>
				<item code="77" name="QCI_77"/>
				<item code="79" name="QCI_79"/>
				<item code="80" name="QCI_80"/>
				<item code="82" name="QCI_82"/>
				<item code="83" name="QCI_83"/>
				<item code="84" name="QCI_84"/>
				<item code="85" name="QCI_85"/>
			</data>
		</avp>

		<avp name="Allocation-Retention-Priority" code="1034" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<!-- 29212 5.3.0 -->
			<data type="Grouped">
                <rule avp="Priority-Level" required="true" max="1"/>
                <rule avp="Pre-emption-Capability" required="false" max="1"/>
                <rule avp="Pre-emption-Vulnerability" required="false" max="1"/>
            </data>
        </avp>

		<avp name="Priority-Level" code="1046" must="V"	may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Pre-emption-Capability" code="1047" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="PRE_EMPTION_CAPABILITY_ENABLED"/>
				<item code="1" name="PRE_EMPTION_CAPABILITY_DISABLED"/>
			</data>
		</avp>

		<avp name="Pre-emption-Vulnerability" code="1048" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="PRE_EMPTION_VULNERABILITY_ENABLED"/>
				<item code="1" name="PRE_EMPTION_VULNERABILITY_DISABLED"/>
			</data>
		</avp>

		<avp name="PDN-GW-Allocation-Type" code="1438" must="M,V" may="-" must-not="-" may-encrypt="-">
			<!-- 29272 7.3.44 -->
			<data type="Enumerated">
				<item code="0" name="STATIC"/>
				<item code="1" name="DYNAMIC"/>
            </data>
        </avp>

		<avp name="VPLMN-Dynamic-Address-Allowed" code="1432" must="M,V" may="-" must-not="-" may-encrypt="-">
			<!-- 29272 7.3.38 -->
			<data type="Enumerated">
				<item code="0" name="NOTALLOWED"/>
				<item code="1" name="ALLOWED"/>
            </data>
        </avp>

		<avp name="AMBR" code="1435" must="M,V" may="-" must-not="-" may-encrypt="N" vendor-id="10415">
			<!-- 29212 7.3.41 -->
			<data type="Grouped">
                <rule avp="Max-Requested-Bandwidth-UL" required="true" max="1"/>
				<rule avp="Max-Requested-Bandwidth-DL" required="true" max="1"/>
            </data>
        </avp>

		<avp name="Max-Requested-Bandwidth-DL" code="515" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Max-Requested-Bandwidth-UL" code="516" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Specific-APN-Info" code="1472" must="M,V" may="-" must-not="-" may-encrypt="N" vendor-id="10415">
			<!-- 29212 7.3.82 -->
			<data type="Grouped">
                <rule avp="Service-Selection" required="true" max="1"/>
				<rule avp="MIP6-Agent-Info" required="true" max="1"/>
				<rule avp="Visited-Network-Identifier" required="false"/>
            </data>
        </avp>

		<avp name="APN-OI-Replacement" code="1427" must="M,V" may="-" may-encrypt="N">
			<!-- 29212 7.3.32 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="SIPTO-Permission" code="1613" must="V" may="-" must-not="M" may-encrypt="N">
			<!-- 29272 7.3.135 -->
			<data type="Enumerated">
				<item code="0" name="SIPTO_above_RAN_ALLOWED"/>
				<item code="1" name="SIPTO_above_RAN_NOTALLOWED"/>
            </data>
        </avp>

		<avp name="LIPA-Permission" code="1618" must="V" may="-" must-not="M" may-encrypt="N">
			<!-- 29272 7.3.133 -->
			<data type="Enumerated">
				<item code="0" name="LIPA_PROHIBITED"/>
				<item code="1" name="LIPA_ONLY"/>
				<item code="2" name="LIPA_CONDITIONAL"/>
            </data>
        </avp>

		<avp name="Restoration-Priority" code="1663" must="V" may="-" must-not="M" may-encrypt="N">
			<!-- 29212 7.3.174 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="SIPTO-Local-Network-Permission" code="1665" must="V" may="-" must-not="M" may-encrypt="N">
			<!-- 29212 7.3.176 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="WLAN-offloadability" code="1667" must="V" may="-" must-not="-" may-encrypt="N">
			<!-- 29212 7.3.181 -->
			<data type="Grouped">
                <rule avp="WLAN-offloadability-EUTRAN" required="false" max="1"/>
				<rule avp="WLAN-offloadability-UTRAN" required="false" max="1"/>
            </data>
        </avp>

		<avp name="WLAN-offloadability-EUTRAN" code="1668" must="V" may="-" must-not="M" may-encrypt="N">
			<!-- 29212 7.3.182 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="WLAN-offloadability-UTRAN" code="1669" must="V" may="-" must-not="M" may-encrypt="N">
			<!-- 29212 7.3.183 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="MIP-Session-Key" code="343" must="M" may="-" must-not="V,P" may-encrypt="N">
			<!-- 29213 9.2.3.6.7 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Time-Quota-Threshold" code="868" must="V,M" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Volume-Quota-Threshold" code="869" must="V,M" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Unit-Quota-Threshold" code="1226" must="V,M" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Quota-Holding-Time" code="871" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Quota-Consumption-Time" code="881" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Reporting-Reason" code="872" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="THRESHOLD"/>
				<item code="1" name="QHT"/>
				<item code="2" name="FINAL"/>
				<item code="3" name="QUOTA_EXHAUSTED"/>
				<item code="4" name="VALIDITY_TIME"/>
				<item code="5" name="OTHER_QUOTA_TYPE"/>
				<item code="6" name="RATING_CONDITION_CHANGE"/>
				<item code="7" name="FORCED_REAUTHORISATION"/>
				<item code="8" name="POOL_EXHAUSTED"/>
				<item code="9" name="UNUSED_QUOTA_TIMER"/>
			</data>
		</avp>

		<avp name="Trigger" code="1264" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Trigger-Type" required="false"/>
			</data>
		</avp>

		<avp name="PS-Furnish-Charging-Information" code="865" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="3GPP-Charging-Id" required="false" max="1"/>
				<rule avp="PS-Free-Format-Data" required="false" max="1"/>
				<rule avp="PS-Append-Free-Format-Data" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Refund-Information" code="2022" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="AF-Correlation-Information" code="1276" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="AF-Charging-Identifier" required="true" max="1"/>
				<rule avp="Flows" required="false"/>
			</data>
		</avp>

		<avp name="Envelope" code="1266" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Envelope-Start-Time" required="true" max="1"/>
				<rule avp="Envelope-End-Time" required="false" max="1"/>
				<rule avp="CC-Total-Octets" required="false" max="1"/>
				<rule avp="CC-Input-Octets" required="false" max="1"/>
				<rule avp="CC-Output-Octets" required="false" max="1"/>
				<rule avp="CC-Service-Specific-Units" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Envelope-Reporting" code="1268" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="DO_NOT_REPORT_ENVELOPES"/>
				<item code="1" name="REPORT_ENVELOPES"/>
				<item code="2" name="REPORT_ENVELOPES_WITH_VOLUME"/>
				<item code="3" name="REPORT_ENVELOPES_WITH_EVENTS"/>
				<item code="4" name="REPORT_ENVELOPES_WITH_VOLUME_AND_EVENTS"/>
			</data>
		</avp>

		<avp name="Time-Quota-Mechanism" code="1270" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Time-Quota-Type" required="true" max="1"/>
				<rule avp="Base-Time-Interval" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Service-Specific-Info" code="1249" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Service-Specific-Data" required="false" max="1"/>
				<rule avp="Service-Specific-Type" required="false" max="1"/>
			</data>
		</avp>

		<avp name="QoS-Information" code="1016" must="V,M" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="Grouped">
				<rule avp="QoS-Class-Identifier" required="false" max="1"/>
				<rule avp="Max-Requested-Bandwidth-UL" required="false" max="1"/>
				<rule avp="Max-Requested-Bandwidth-DL" required="false" max="1"/>
				<rule avp="Extended-Max-Requested-BW-UL" required="false" max="1"/>
				<rule avp="Extended-Max-Requested-BW-DL" required="false" max="1"/>
				<rule avp="Guaranteed-Bitrate-UL" required="false" max="1"/>
				<rule avp="Guaranteed-Bitrate-DL" required="false" max="1"/>
				<rule avp="Extended-GBR-UL" required="false" max="1"/>
				<rule avp="Extended-GBR-DL" required="false" max="1"/>
				<rule avp="Bearer-Identifier" required="false" max="1"/>
				<rule avp="Allocation-Retention-Priority" required="false" max="1"/>
				<rule avp="APN-Aggregate-Max-Bitrate-UL" required="false" max="1"/>
				<rule avp="APN-Aggregate-Max-Bitrate-DL" required="false" max="1"/>
				<rule avp="Extended-APN-AMBR-UL" required="false" max="1"/>
				<rule avp="Extended-APN-AMBR-DL" required="false" max="1"/>
				<rule avp="Conditional-APN-Aggregate-Max-Bitrate" required="false"/>
			</data>
		</avp>

		<avp name="Announcement-Information" code="3904" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Announcement-Identifier" required="true" max="1"/>
				<rule avp="Variable-Part" required="false"/>
				<rule avp="Time-Indicator" required="false" max="1"/>
				<rule avp="Quota-Indicator" required="false" max="1"/>
				<rule avp="Announcement-Order" required="false" max="1"/>
				<rule avp="Play-Alternative" required="false" max="1"/>
				<rule avp="Privacy-Indicator" required="false" max="1"/>
				<rule avp="Language" required="false" max="1"/>
			</data>
		</avp>

		<avp name="3GPP-RAT-Type" code="21" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Related-Trigger" code="3926" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Trigger-Type" required="false"/>
			</data>
		</avp>

		<avp name="Trigger-Type" code="870" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="1" name="CHANGE_IN_SGSN_IP_ADDRESS"/>
				<item code="2" name="CHANGE_IN_QOS"/>
				<item code="3" name="CHANGE_IN_LOCATION"/>
				<item code="4" name="CHANGE_IN_RAT"/>
				<item code="5" name="CHANGE_IN_UE_TIMEZONE"/>
				<item code="10" name="CHANGEINQOS_TRAFFIC_CLASS"/>
				<item code="11" name="CHANGEINQOS_RELIABILITY_CLASS"/>
				<item code="12" name="CHANGEINQOS_DELAY_CLASS"/>
				<item code="13" name="CHANGEINQOS_PEAK_THROUGHPUT"/>
				<item code="14" name="CHANGEINQOS_PRECEDENCE_CLASS"/>
				<item code="15" name="CHANGEINQOS_MEAN_THROUGHPUT"/>
				<item code="16" name="CHANGEINQOS_MAXIMUM_BIT_RATE_FOR_UPLINK"/>
				<item code="17" name="CHANGEINQOS_MAXIMUM_BIT_RATE_FOR_DOWNLINK"/>
				<item code="18" name="CHANGEINQOS_RESIDUAL_BER"/>
				<item code="19" name="CHANGEINQOS_SDU_ERROR_RATIO"/>
				<item code="20" name="CHANGEINQOS_TRANSFER_DELAY"/>
				<item code="21" name="CHANGEINQOS_TRAFFIC_HANDLING_PRIORITY"/>
				<item code="22" name="CHANGEINQOS_GUARANTEED_BIT_RATE_FOR_UPLINK"/>
				<item code="23" name="CHANGEINQOS_GUARANTEED_BIT_RATE_FOR_DOWNLINK"/>
				<item code="24" name="CHANGEINQOS_APN_AGGREGATE_MAXIMUM_BIT_RATE"/>
				<item code="30" name="CHANGEINLOCATION_MCC"/>
				<item code="31" name="CHANGEINLOCATION_MNC"/>
				<item code="32" name="CHANGEINLOCATION_RAC"/>
				<item code="33" name="CHANGEINLOCATION_LAC"/>
				<item code="34" name="CHANGEINLOCATION_CELLID"/>
				<item code="35" name="CHANGEINLOCATION_TAC"/>
				<item code="36" name="CHANGEINLOCATION_ECGI"/>
				<item code="40" name="CHANGE_IN_MEDIA_COMPOSITION"/>
				<item code="50" name="CHANGE_IN_PARTICIPANTS_NMB"/>
				<item code="51" name="CHANGE_IN__THRSHLD_OF_PARTICIPANTS_NMB"/>
				<item code="52" name="CHANGE_IN_USER_PARTICIPATING_TYPE"/>
				<item code="60" name="CHANGE_IN_SERVICE_CONDITION"/>
				<item code="61" name="CHANGE_IN_SERVING_NODE"/>
				<item code="62" name="CHANGE_IN_ACCESS_FOR_A_SERVICE_DATA_FLOW"/>
				<item code="70" name="CHANGE_IN_USER_CSG_INFORMATION"/>
				<item code="71" name="CHANGE_IN_HYBRID_SUBSCRIBED_USER_CSG_INFORMATION"/>
				<item code="72" name="CHANGE_IN_HYBRID_UNSUBSCRIBED_USER_CSG_INFORMATION"/>
				<item code="73" name="CHANGE_OF_UE_PRESENCE_IN_PRESENCE_REPORTING_AREA"/>
				<item code="75" name="CHANGE_IN_APN_RATE_CONTROL"/>
				<item code="76" name="CHANGE_IN_3GPP_PS_DATA_OFF"/>
			</data>
		</avp>

		<avp name="Envelope-Start-Time" code="1269" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Envelope-End-Time" code="1267" must="V,M" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Time-Quota-Type" code="1271" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="DISCRETE_TIME_PERIOD"/>
				<item code="1" name="CONTINUOUS_TIME_PERIOD"/>
			</data>
		</avp>

		<avp name="Base-Time-Interval" code="1265" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Service-Specific-Data" code="863" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Service-Specific-Type" code="1257" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Announcement-Identifier" code="3905" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Variable-Part" code="3907" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Variable-Part-Order" required="false" max="1"/>
				<rule avp="Variable-Part-Type" required="true" max="1"/>
				<rule avp="Variable-Part-Value" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Time-Indicator" code="3911" must="V,M" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Quota-Indicator" code="3912" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="QUOTA_IS_NOT_USED_DURING_PLAYBACK"/>
				<item code="1" name="QUOTA_IS_USED_DURING_PLAYBACK"/>
			</data>
		</avp>

		<avp name="Announcement-Order" code="3906" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Play-Alternative" code="3913" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="SERVED_PARTY"/>
				<item code="1" name="REMOTE_PARTY"/>
			</data>
		</avp>

		<avp name="Privacy-Indicator" code="3915" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NOT_PRIVATE"/>
				<item code="1" name="PRIVATE"/>
			</data>
		</avp>

		<avp name="Language" code="3914" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Variable-Part-Order" code="3908" must="V,M" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Variable-Part-Type" code="3909" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Variable-Part-Value" code="3910" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="AoC-Information" code="2054" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="AoC-Cost-Information" required="false" max="1"/>
				<rule avp="Tariff-Information" required="false" max="1"/>
				<rule avp="AoC-Subscription-Information" required="false" max="1"/>
			</data>
		</avp>

		<avp name="PS-Information" code="874" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Supported-Features" required="false"/>
				<rule avp="3GPP-Charging-Id" required="false" max="1"/>
				<rule avp="3GPP-GPRS-Negotiated-QoS-Profile" required="false" max="1"/>
				<rule avp="Node-Id" required="false" max="1"/>
				<rule avp="3GPP-PDP-Type" required="false" max="1"/>
				<rule avp="PDN-Connection-Charging-ID" required="false" max="1"/>
				<rule avp="PDP-Address" required="false"/>
				<rule avp="PDP-Address-Prefix-Length" required="false" max="1"/>
				<rule avp="Dynamic-Address-Flag" required="false" max="1"/>
				<rule avp="Dynamic-Address-Flag-Extension" required="false" max="1"/>
				<rule avp="QoS-Information" required="false" max="1"/>
				<rule avp="SGSN-Address" required="false"/>
				<rule avp="GGSN-Address" required="false"/>
				<rule avp="TDF-IP-Address" required="false"/>
				<rule avp="SGW-Address" required="false"/>
				<rule avp="EPDG-Address" required="false"/>
				<rule avp="TWAG-Address" required="false"/>
				<rule avp="CG-Address" required="false" max="1"/>
				<rule avp="Serving-Node-Type" required="false" max="1"/>
				<rule avp="SGW-Change" required="false" max="1"/>
				<rule avp="3GPP-IMSI-MCC-MNC" required="false" max="1"/>
				<rule avp="IMSI-Unauthenticated-Flag" required="false" max="1"/>
				<rule avp="3GPP-GGSN-MCC-MNC" required="false" max="1"/>
				<rule avp="3GPP-NSAPI" required="false" max="1"/>
				<rule avp="Called-Station-Id" required="false" max="1"/>
				<rule avp="3GPP-Session-Stop-Indicator" required="false" max="1"/>
				<rule avp="3GPP-Selection-Mode" required="false" max="1"/>
				<rule avp="3GPP-Charging-Characteristics" required="false" max="1"/>
				<rule avp="Charging-Characteristics-Selection-Mode" required="false" max="1"/>
				<rule avp="3GPP-SGSN-MCC-MNC" required="false" max="1"/>
				<rule avp="3GPP-MS-TimeZone" required="false" max="1"/>
				<rule avp="Charging-Rule-Base-Name" required="false" max="1"/>
				<rule avp="ADC-Rule-Base-Name" required="false" max="1"/>
				<rule avp="3GPP-User-Location-Info" required="false" max="1"/>
				<rule avp="User-Location-Info-Time" required="false" max="1"/>
				<rule avp="User-CSG-Information" required="false" max="1"/>
				<rule avp="Presence-Reporting-Area-Information" required="false"/>
				<rule avp="3GPP2-BSID" required="false" max="1"/>
				<rule avp="TWAN-User-Location-Info" required="false" max="1"/>
				<rule avp="UWAN-User-Location-Info" required="false" max="1"/>
				<rule avp="3GPP-RAT-Type" required="false" max="1"/>
				<rule avp="PS-Furnish-Charging-Information" required="false" max="1"/>
				<rule avp="PDP-Context-Type" required="false" max="1"/>
				<rule avp="Offline-Charging" required="false" max="1"/>
				<rule avp="Traffic-Data-Volumes" required="false"/>
				<rule avp="Service-Data-Container" required="false"/>
				<rule avp="User-Equipment-Info" required="false" max="1"/>
				<rule avp="Terminal-Information" required="false" max="1"/>
				<rule avp="Start-Time" required="false" max="1"/>
				<rule avp="Stop-Time" required="false" max="1"/>
				<rule avp="Change-Condition" required="false" max="1"/>
				<rule avp="Diagnostics" required="false" max="1"/>
				<rule avp="Low-Priority-Indicator" required="false" max="1"/>
				<rule avp="NBIFOM-Mode" required="false" max="1"/>
				<rule avp="NBIFOM-Support" required="false" max="1"/>
				<rule avp="MME-Number-for-MT-SMS" required="false" max="1"/>
				<rule avp="MME-Name" required="false" max="1"/>
				<rule avp="MME-Realm" required="false" max="1"/>
				<rule avp="Logical-Access-ID" required="false" max="1"/>
				<rule avp="Physical-Access-ID" required="false" max="1"/>
				<rule avp="Fixed-User-Location-Info" required="false" max="1"/>
				<rule avp="CN-Operator-Selection-Entity" required="false" max="1"/>
				<rule avp="Enhanced-Diagnostics" required="false" max="1"/>
				<rule avp="SGi-PtP-Tunnelling-Method" required="false" max="1"/>
				<rule avp="CP-CIoT-EPS-Optimisation-Indicator" required="false" max="1"/>
				<rule avp="UNI-PDU-CP-Only-Flag" required="false" max="1"/>
				<rule avp="Serving-PLMN-Rate-Control" required="false" max="1"/>
				<rule avp="APN-Rate-Control" required="false" max="1"/>
				<rule avp="Charging-Per-IP-CAN-Session-Indicator" required="false" max="1"/>
				<rule avp="RRC-Cause-Counter" required="false" max="1"/>
				<rule avp="3GPP-PS-Data-Off-Status" required="false" max="1"/>
				<rule avp="SCS-AS-Address" required="false" max="1"/>
				<rule avp="Unused-Quota-Timer" required="false" max="1"/>
				<rule avp="RAN-Secondary-RAT-Usage-Report" required="false"/>
			</data>
		</avp>

		<avp name="IMS-Information" code="876" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Event-Type" required="false" max="1"/>
				<rule avp="Role-Of-Node" required="false" max="1"/>
				<rule avp="Node-Functionality" required="true" max="1"/>
				<rule avp="User-Session-Id" required="false" max="1"/>
				<rule avp="Outgoing-Session-Id" required="false" max="1"/>
				<rule avp="Session-Priority" required="false" max="1"/>
				<rule avp="Calling-Party-Address" required="false"/>
				<rule avp="Called-Party-Address" required="false" max="1"/>
				<rule avp="Called-Asserted-Identity" required="false"/>
				<rule avp="Called-Identity-Change" required="false" max="1"/>
				<rule avp="Number-Portability-Routing-Information" required="false" max="1"/>
				<rule avp="Carrier-Select-Routing-Information" required="false" max="1"/>
				<rule avp="Alternate-Charged-Party-Address" required="false" max="1"/>
				<rule avp="Requested-Party-Address" required="false"/>
				<rule avp="Associated-URI" required="false"/>
				<rule avp="Time-Stamps" required="false" max="1"/>
				<rule avp="Application-Server-Information" required="false"/>
				<rule avp="Inter-Operator-Identifier" required="false"/>
				<rule avp="Transit-IOI-List" required="false"/>
				<rule avp="IMS-Charging-Identifier" required="false" max="1"/>
				<rule avp="SDP-Session-Description" required="false"/>
				<rule avp="SDP-Media-Component" required="false"/>
				<rule avp="Served-Party-IP-Address" required="false" max="1"/>
				<rule avp="Server-Capabilities" required="false" max="1"/>
				<rule avp="Trunk-Group-ID" required="false" max="1"/>
				<rule avp="Bearer-Service" required="false" max="1"/>
				<rule avp="Service-Id" required="false" max="1"/>
				<rule avp="Service-Specific-Info" required="false"/>
				<rule avp="Message-Body" required="false"/>
				<rule avp="Cause-Code" required="false" max="1"/>
				<rule avp="Reason-Header" required="false"/>
				<rule avp="Access-Network-Information" required="false"/>
				<rule avp="Cellular-Network-Information" required="false" max="1"/>
				<rule avp="Early-Media-Description" required="false"/>
				<rule avp="IMS-Communication-Service-Identifier" required="false" max="1"/>
				<rule avp="IMS-Application-Reference-Identifier" required="false" max="1"/>
				<rule avp="Online-Charging-Flag" required="false" max="1"/>
				<rule avp="Real-Time-Tariff-Information" required="false" max="1"/>
				<rule avp="Account-Expiration" required="false" max="1"/>
				<rule avp="Initial-IMS-Charging-Identifier" required="false" max="1"/>
				<rule avp="NNI-Information" required="false"/>
				<rule avp="From-Address" required="false" max="1"/>
				<rule avp="IMS-Emergency-Indicator" required="false" max="1"/>
				<rule avp="IMS-Visited-Network-Identifier" required="false" max="1"/>
				<rule avp="Access-Network-Info-Change" required="false"/>
				<rule avp="Access-Transfer-Information" required="false"/>
				<rule avp="Related-IMS-Charging-Identifier" required="false" max="1"/>
				<rule avp="Related-IMS-Charging-Identifier-Node" required="false" max="1"/>
				<rule avp="Route-Header-Received" required="false" max="1"/>
				<rule avp="Route-Header-Transmitted" required="false" max="1"/>
				<rule avp="Instance-Id" required="false" max="1"/>
				<rule avp="TAD-Identifier" required="false" max="1"/>
				<rule avp="FE-Identifier-List" required="false" max="1"/>
			</data>
		</avp>

		<avp name="MMS-Information" code="877" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Originator-Address" required="false" max="1"/>
				<rule avp="Recipient-Address" required="false"/>
				<rule avp="Submission-Time" required="false" max="1"/>
				<rule avp="MM-Content-Type" required="false" max="1"/>
				<rule avp="Priority" required="false" max="1"/>
				<rule avp="Message-ID" required="false" max="1"/>
				<rule avp="Message-Type" required="false" max="1"/>
				<rule avp="Message-Size" required="false" max="1"/>
				<rule avp="Message-Class" required="false" max="1"/>
				<rule avp="Delivery-Report-Requested" required="false" max="1"/>
				<rule avp="Read-Reply-Report-Requested" required="false" max="1"/>
				<rule avp="MMBox-Storage-Requested" required="false" max="1"/>
				<rule avp="Applic-ID" required="false" max="1"/>
				<rule avp="Reply-Applic-ID" required="false" max="1"/>
				<rule avp="Aux-Applic-Info" required="false" max="1"/>
				<rule avp="Content-Class" required="false" max="1"/>
				<rule avp="DRM-Content" required="false" max="1"/>
				<rule avp="Adaptations" required="false" max="1"/>
				<rule avp="VASP-Id" required="false" max="1"/>
				<rule avp="VAS-Id" required="false" max="1"/>
			</data>
		</avp>

		<avp name="LCS-Information" code="878" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="LCS-Client-ID" required="false" max="1"/>
				<rule avp="Location-Type" required="false" max="1"/>
				<rule avp="Location-Estimate" required="false" max="1"/>
				<rule avp="Positioning-Data" required="false" max="1"/>
				<rule avp="3GPP-IMSI" required="false" max="1"/>
				<rule avp="MSISDN" required="false" max="1"/>
			</data>
		</avp>

		<avp name="PoC-Information" code="879" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="PoC-Server-Role" required="false" max="1"/>
				<rule avp="PoC-Session-Type" required="false" max="1"/>
				<rule avp="PoC-User-Role" required="false" max="1"/>
				<rule avp="PoC-Session-Initiation-Type" required="false" max="1"/>
				<rule avp="PoC-Event-Type" required="false" max="1"/>
				<rule avp="Number-Of-Participants" required="false" max="1"/>
				<rule avp="Participants-Involved" required="false"/>
				<rule avp="Participant-Group" required="false"/>
				<rule avp="Talk-Burst-Exchange" required="false"/>
				<rule avp="PoC-Controlling-Address" required="false" max="1"/>
				<rule avp="PoC-Group-Name" required="false" max="1"/>
				<rule avp="PoC-Session-Id" required="false" max="1"/>
				<rule avp="Charged-Party" required="false" max="1"/>
			</data>
		</avp>

		<avp name="MBMS-Information" code="880" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="TMGI" required="false" max="1"/>
				<rule avp="MBMS-Service-Type" required="false" max="1"/>
				<rule avp="MBMS-User-Service-Type" required="false" max="1"/>
				<rule avp="File-Repair-Supported" required="false" max="1"/>
				<rule avp="Required-MBMS-Bearer-Capabilities" required="false" max="1"/>
				<rule avp="MBMS-2G-3G-Indicator" required="false" max="1"/>
				<rule avp="RAI" required="false" max="1"/>
				<rule avp="MBMS-Service-Area" required="false"/>
				<rule avp="MBMS-Session-Identity" required="false" max="1"/>
				<rule avp="CN-IP-Multicast-Distribution" required="false" max="1"/>
				<rule avp="MBMS-GW-Address" required="false" max="1"/>
				<rule avp="MBMS-Charged-Party" required="false" max="1"/>
				<rule avp="MSISDN" required="false"/>
				<rule avp="MBMS-Data-Transfer-Start" required="false" max="1"/>
				<rule avp="MBMS-Data-Transfer-Stop" required="false" max="1"/>
			</data>
		</avp>

		<avp name="SMS-Information" code="2000" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SMS-Node" required="false" max="1"/>
				<rule avp="Client-Address" required="false" max="1"/>
				<rule avp="Originator-SCCP-Address" required="false" max="1"/>
				<rule avp="SMSC-Address" required="false" max="1"/>
				<rule avp="Data-Coding-Scheme" required="false" max="1"/>
				<rule avp="SM-Discharge-Time" required="false" max="1"/>
				<rule avp="SM-Message-Type" required="false" max="1"/>
				<rule avp="Originator-Interface" required="false" max="1"/>
				<rule avp="SM-Protocol-ID" required="false" max="1"/>
				<rule avp="Reply-Path-Requested" required="false" max="1"/>
				<rule avp="SM-Status" required="false" max="1"/>
				<rule avp="SM-User-Data-Header" required="false" max="1"/>
				<rule avp="Number-Of-Messages-Sent" required="false" max="1"/>
				<rule avp="SM-Sequence-Number" required="false" max="1"/>
				<rule avp="Recipient-Info" required="false"/>
				<rule avp="Originator-Received-Address" required="false" max="1"/>
				<rule avp="SM-Service-Type" required="false" max="1"/>
				<rule avp="SMS-Result" required="false" max="1"/>
				<rule avp="SM-Device-Trigger-Indicator" required="false" max="1"/>
				<rule avp="SM-Device-Trigger-Information" required="false" max="1"/>
				<rule avp="MTC-IWF-Address" required="false" max="1"/>
				<rule avp="Application-Port-Identifier" required="false" max="1"/>
				<rule avp="External-Identifier" required="false" max="1"/>
			</data>
		</avp>

		<avp name="VCS-Information" code="3410" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Bearer-Capability" required="false" max="1"/>
				<rule avp="Network-Call-Reference-Number" required="false" max="1"/>
				<rule avp="MSC-Address" required="false" max="1"/>
				<rule avp="Basic-Service-Code" required="false" max="1"/>
				<rule avp="ISUP-Location-Number" required="false" max="1"/>
				<rule avp="VLR-Number" required="false" max="1"/>
				<rule avp="Forwarding-Pending" required="false" max="1"/>
				<rule avp="ISUP-Cause" required="false" max="1"/>
				<rule avp="Start-Time" required="false" max="1"/>
				<rule avp="Start-of-Charging" required="false" max="1"/>
				<rule avp="Stop-Time" required="false" max="1"/>
				<rule avp="PS-Free-Format-Data" required="false" max="1"/>
			</data>
		</avp>

		<avp name="MMTel-Information" code="2030" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Supplementary-Service" required="false"/>
			</data>
		</avp>

		<avp name="ProSe-Information" code="3447" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Supported-Features" required="false"/>
				<rule avp="Announcing-UE-HPLMN-Identifier" required="false" max="1"/>
				<rule avp="Announcing-UE-VPLMN-Identifier" required="false" max="1"/>
				<rule avp="Monitoring-UE-HPLMN-Identifier" required="false" max="1"/>
				<rule avp="Monitoring-UE-VPLMN-Identifier" required="false" max="1"/>
				<rule avp="Monitored-PLMN-Identifier" required="false" max="1"/>
				<rule avp="Role-Of-ProSe-Function" required="false" max="1"/>
				<rule avp="ProSe-App-Id" required="false" max="1"/>
				<rule avp="ProSe-3rd-Party-Application-ID" required="false" max="1"/>
				<rule avp="Application-Specific-Data" required="false" max="1"/>
				<rule avp="ProSe-Event-Type" required="false" max="1"/>
				<rule avp="ProSe-Direct-Discovery-Model" required="false" max="1"/>
				<rule avp="ProSe-Function-IP-Address" required="false" max="1"/>
				<rule avp="ProSe-Function-ID" required="false" max="1"/>
				<rule avp="ProSe-Validity-Timer" required="false" max="1"/>
				<rule avp="ProSe-Role-Of-UE" required="false" max="1"/>
				<rule avp="ProSe-Request-Timestamp" required="false" max="1"/>
				<rule avp="PC3-Control-Protocol-Cause" required="false" max="1"/>
				<rule avp="Monitoring-UE-Identifier" required="false" max="1"/>
				<rule avp="ProSe-Function-PLMN-Identifier" required="false" max="1"/>
				<rule avp="Requestor-PLMN-Identifier" required="false" max="1"/>
				<rule avp="Origin-App-Layer-User-Id" required="false" max="1"/>
				<rule avp="WLAN-Link-Layer-Id" required="false" max="1"/>
				<rule avp="Requesting-EPUID" required="false" max="1"/>
				<rule avp="Target-App-Layer-User-Id" required="false" max="1"/>
				<rule avp="Requested-PLMN-Identifier" required="false" max="1"/>
				<rule avp="Time-Window" required="false" max="1"/>
				<rule avp="ProSe-Range-Class" required="false" max="1"/>
				<rule avp="Proximity-Alert-Indication" required="false" max="1"/>
				<rule avp="Proximity-Alert-Timestamp" required="false" max="1"/>
				<rule avp="Proximity-Cancellation-Timestamp" required="false" max="1"/>
				<rule avp="ProSe-Reason-For-Cancellation" required="false" max="1"/>
				<rule avp="PC3-EPC-Control-Protocol-Cause" required="false" max="1"/>
				<rule avp="ProSe-UE-ID" required="false" max="1"/>
				<rule avp="ProSe-Source-IP-Address" required="false" max="1"/>
				<rule avp="Layer-2-Group-ID" required="false" max="1"/>
				<rule avp="ProSe-Group-IP-Multicast-Address" required="false" max="1"/>
				<rule avp="Coverage-Info" required="false"/>
				<rule avp="Radio-Parameter-Set-Info" required="false"/>
				<rule avp="Transmitter-Info" required="false"/>
				<rule avp="Time-First-Transmission" required="false" max="1"/>
				<rule avp="Time-First-Reception" required="false" max="1"/>
				<rule avp="ProSe-Direct-Communication-Transmission-Data-Container" required="false"/>
				<rule avp="ProSe-Direct-Communication-Reception-Data-Container" required="false"/>
				<rule avp="Announcing-PLMN-ID" required="false" max="1"/>
				<rule avp="ProSe-Target-Layer-2-ID" required="false" max="1"/>
				<rule avp="Relay-IP-address" required="false" max="1"/>
				<rule avp="ProSe-UE-to-Network-Relay-UE-ID" required="false" max="1"/>
				<rule avp="Target-IP-Address" required="false" max="1"/>
				<rule avp="PC5-Radio-Technology" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Service-Generic-Information" code="1256" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Application-Server-ID" required="false" max="1"/>
				<rule avp="Application-Service-Type" required="false" max="1"/>
				<rule avp="Application-Session-ID" required="false" max="1"/>
				<rule avp="Delivery-Status" required="false" max="1"/>
			</data>
		</avp>

		<avp name="IM-Information" code="2110" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Total-Number-Of-Messages-Sent" required="false" max="1"/>
				<rule avp="Total-Number-Of-Messages-Exploded" required="false" max="1"/>
				<rule avp="Number-Of-Messages-Successfully-Sent" required="false" max="1"/>
				<rule avp="Number-Of-Messages-Successfully-Exploded" required="false" max="1"/>
			</data>
		</avp>

		<avp name="DCD-Information" code="2115" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Content-ID" required="false" max="1"/>
				<rule avp="Content-Provider-ID" required="false" max="1"/>
			</data>
		</avp>

		<avp name="M2M-Information" code="1011" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Grouped">
				<rule avp="Application-Entity-ID" required="false" max="1"/>
				<rule avp="External-ID" required="false" max="1"/>
				<rule avp="Receiver" required="false" max="1"/>
				<rule avp="Originator" required="false" max="1"/>
				<rule avp="Hosting-CSE-ID" required="false" max="1"/>
				<rule avp="Target-ID" required="false" max="1"/>
				<rule avp="Protocol-Type" required="false" max="1"/>
				<rule avp="Request-Operation" required="false" max="1"/>
				<rule avp="Request-Headers-Size" required="false" max="1"/>
				<rule avp="Request-Body-Size" required="false" max="1"/>
				<rule avp="Response-Headers-Size" required="false" max="1"/>
				<rule avp="Response-Body-Size" required="false" max="1"/>
				<rule avp="Response-Status-Code" required="false" max="1"/>
				<rule avp="Rating-Group" required="false" max="1"/>
				<rule avp="M2M-Event-Record-Timestamp" required="false" max="1"/>
				<rule avp="Control-Memory-Size" required="false" max="1"/>
				<rule avp="Data-Memory-Size" required="false" max="1"/>
				<rule avp="Access-Network-Identifier" required="false" max="1"/>
				<rule avp="Occupancy" required="false" max="1"/>
				<rule avp="Group-Name" required="false" max="1"/>
				<rule avp="Maximum-Number-Members" required="false" max="1"/>
				<rule avp="Current-Number-Members" required="false" max="1"/>
				<rule avp="Subgroup-Name" required="false" max="1"/>
				<rule avp="Node-Id" required="false" max="1"/>
			</data>
		</avp>

		<avp name="CPDT-Information" code="3927" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="External-Identifier" required="false" max="1"/>
				<rule avp="SCEF-ID" required="false" max="1"/>
				<rule avp="Serving-Node-Identity" required="false" max="1"/>
				<rule avp="SGW-Change" required="false" max="1"/>
				<rule avp="NIDD-Submission" required="false" max="1"/>
			</data>
		</avp>

		<avp name="AoC-Cost-Information" code="2053" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Accumulated-Cost" required="false" max="1"/>
				<rule avp="Incremental-Cost" required="false"/>
				<rule avp="Currency-Code" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Tariff-Information" code="2060" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Current-Tariff" required="true" max="1"/>
				<rule avp="Tariff-Time-Change" required="false" max="1"/>
				<rule avp="Next-Tariff" required="false" max="1"/>
			</data>
		</avp>

		<avp name="AoC-Subscription-Information" code="2314" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="AoC-Service" required="false"/>
				<rule avp="AoC-Format" required="false" max="1"/>
				<rule avp="Preferred-AoC-Currency" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Accumulated-Cost" code="2052" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Value-Digits" required="true" max="1"/>
				<rule avp="Exponent" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Incremental-Cost" code="2062" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Value-Digits" required="true" max="1"/>
				<rule avp="Exponent" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Current-Tariff" code="2056" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Currency-Code" required="false" max="1"/>
				<rule avp="Scale-Factor" required="false" max="1"/>
				<rule avp="Rate-Element" required="false"/>
			</data>
		</avp>

		<avp name="Next-Tariff" code="2057" must="V/M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Currency-Code" required="false" max="1"/>
				<rule avp="Scale-Factor" required="false" max="1"/>
				<rule avp="Rate-Element" required="false"/>
			</data>
		</avp>

		<avp name="Scale-Factor" code="2059" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Value-Digits" required="true" max="1"/>
				<rule avp="Exponent" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Rate-Element" code="2058" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="CC-Unit-Type" required="true" max="1"/>
				<rule avp="Charge-Reason-Code" required="false" max="1"/>
				<rule avp="Unit-Value" required="false" max="1"/>
				<rule avp="Unit-Cost" required="false" max="1"/>
				<rule avp="Unit-Quota-Threshold" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Charge-Reason-Code" code="2118" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="UNKNOWN"/>
				<item code="1" name="USAGE"/>
				<item code="2" name="COMMUNICATION_ATTEMPT_CHARGE"/>
				<item code="3" name="SETUP_CHARGE"/>
				<item code="4" name="ADD_ON_CHARGE"/>
			</data>
		</avp>

		<avp name="Unit-Cost" code="2061" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Value-Digits" required="true" max="1"/>
				<rule avp="Exponent" required="false" max="1"/>
			</data>
		</avp>

		<avp name="AoC-Service" code="2311" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="AoC-Service-Obligatory-Type" required="false" max="1"/>
				<rule avp="AoC-Service-Type" required="false" max="1"/>
			</data>
		</avp>

		<avp name="AoC-Format" code="2310" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="MONETARY"/>
				<item code="1" name="NON_MONETARY"/>
				<item code="2" name="CAI"/>
			</data>
		</avp>

		<avp name="Preferred-AoC-Currency" code="2315" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="AoC-Service-Obligatory-Type" code="2312" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NON_BINDING"/>
				<item code="1" name="BINDING"/>
			</data>
		</avp>

		<avp name="AoC-Service-Type" code="2313" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NONE"/>
				<item code="1" name="AOC_S"/>
				<item code="2" name="AOC_D"/>
				<item code="3" name="AOC_E"/>
			</data>
		</avp>

		<avp name="3GPP-Charging-Id" code="2" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="PDN-Connection-Charging-ID" code="2050" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Node-Id" code="2064" must="V,M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="PDP-Address" code="1227" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="PDP-Address-Prefix-Length" code="2606" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Dynamic-Address-Flag" code="2051" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="Static"/>
				<item code="1" name="Dynamic"/>
			</data>
		</avp>

		<avp name="Dynamic-Address-Flag-Extension" code="2068" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="Static"/>
				<item code="1" name="Dynamic"/>
			</data>
		</avp>

		<avp name="SGSN-Address" code="1228" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="GGSN-Address" code="847" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="TDF-IP-Address" code="1091" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="SGW-Address" code="2067" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="EPDG-Address" code="3425" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="TWAG-Address" code="3903" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="CG-Address" code="846" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Serving-Node-Type" code="2047" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="SGSN"/>
				<item code="1" name="PMIPSGW"/>
				<item code="2" name="GTPSGW"/>
				<item code="3" name="EPDG"/>
				<item code="4" name="HSGW"/>
				<item code="5" name="MME"/>
				<item code="6" name="TWAN"/>
			</data>
		</avp>

		<avp name="SGW-Change" code="2065" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="ACR_START_NOT_DUE_TO_SGW_CHANGE"/>
				<item code="1" name="ACR_START_DUE_TO_SGW_CHANGE"/>
			</data>
		</avp>

		<avp name="3GPP-GPRS-Negotiated-QoS-Profile" code="5" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="3GPP-IMSI-MCC-MNC" code="8" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="IMSI-Unauthenticated-Flag" code="2308" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="AUTHENTICATED"/>
				<item code="1" name="UNAUTHENTICATED"/>
			</data>
		</avp>

		<avp name="3GPP-GGSN-MCC-MNC" code="9" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="3GPP-NSAPI" code="10" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="3GPP-Session-Stop-Indicator" code="11" must="V" may="P" must-not="M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="3GPP-Selection-Mode" code="12" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Charging-Characteristics-Selection-Mode" code="2066" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="SERVING_NODE_SUPPLIED"/>
				<item code="1" name="SUBSCRIPTION_SPECIFIC"/>
				<item code="2" name="APN_SPECIFIC"/>
				<item code="3" name="HOME_DEFAULT"/>
				<item code="4" name="ROAMING_DEFAULT"/>
				<item code="5" name="VISITING_DEFAULT"/>
			</data>
		</avp>

		<avp name="Charging-Rule-Base-Name" code="1004" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="ADC-Rule-Base-Name" code="1095" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="User-CSG-Information" code="2319" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="CSG-Id" required="true" max="1"/>
				<rule avp="CSG-Access-Mode" required="true" max="1"/>
				<rule avp="CSG-Membership-Indication" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Presence-Reporting-Area-Information" code="2822" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Presence-Reporting-Area-Identifier" required="false" max="1"/>
				<rule avp="Presence-Reporting-Area-Status" required="false" max="1"/>
				<rule avp="Presence-Reporting-Area-Elements-List" required="false" max="1"/>
				<rule avp="Presence-Reporting-Area-Node" required="false" max="1"/>
			</data>
		</avp>

		<avp name="3GPP2-BSID" code="9010" must="M,V" may="P" may-encrypt="No" vendor-id="5535">
			<data type="UTF8String"/>
		</avp>

		<avp name="TWAN-User-Location-Info" code="2714" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SSID" required="true" max="1"/>
				<rule avp="BSSID" required="false" max="1"/>
				<rule avp="Civic-Address-Information" required="false" max="1"/>
				<rule avp="WLAN-Operator-Id" required="false" max="1"/>
				<rule avp="Logical-Access-ID" required="false" max="1"/>
			</data>
		</avp>

		<avp name="UWAN-User-Location-Info" code="3918" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="UE-Local-IP-Address" required="true" max="1"/>
				<rule avp="UDP-Source-Port" required="false" max="1"/>
				<rule avp="SSID" required="false" max="1"/>
				<rule avp="BSSID" required="false" max="1"/>
				<rule avp="TCP-Source-Port" required="false" max="1"/>
				<rule avp="Civic-Address-Information" required="false" max="1"/>
				<rule avp="WLAN-Operator-Id" required="false" max="1"/>
				<rule avp="Logical-Access-ID" required="false" max="1"/>
			</data>
		</avp>

		<avp name="PDP-Context-Type" code="1247" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="PRIMARY"/>
				<item code="1" name="SECONDARY"/>
			</data>
		</avp>

		<avp name="Offline-Charging" code="1278" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Quota-Consumption-Time" required="false" max="1"/>
				<rule avp="Time-Quota-Mechanism" required="false" max="1"/>
				<rule avp="Envelope-Reporting" required="false" max="1"/>
				<rule avp="Multiple-Services-Credit-Control" required="false"/>
			</data>
		</avp>

		<avp name="Traffic-Data-Volumes" code="2046" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="QoS-Information" required="false" max="1"/>
				<rule avp="Accounting-Input-Octets" required="false" max="1"/>
				<rule avp="Accounting-Output-Octets" required="false" max="1"/>
				<rule avp="Change-Condition" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
				<rule avp="3GPP-User-Location-Info" required="false" max="1"/>
				<rule avp="UWAN-User-Location-Info" required="false" max="1"/>
				<rule avp="3GPP-Charging-Id" required="false" max="1"/>
				<rule avp="Presence-Reporting-Area-Status" required="false" max="1"/>
				<rule avp="Presence-Reporting-Area-Information" required="false"/>
				<rule avp="User-CSG-Information" required="false" max="1"/>
				<rule avp="3GPP-RAT-Type" required="false" max="1"/>
				<rule avp="Access-Availability-Change-Reason" required="false" max="1"/>
				<rule avp="Related-Change-Condition-Information" required="false" max="1"/>
				<rule avp="Diagnostics" required="false" max="1"/>
				<rule avp="Enhanced-Diagnostics" required="false" max="1"/>
				<rule avp="CP-CIoT-EPS-Optimisation-Indicator" required="false" max="1"/>
				<rule avp="Serving-PLMN-Rate-Control" required="false" max="1"/>
				<rule avp="APN-Rate-Control" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Service-Data-Container" code="2040" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="AF-Correlation-Information" required="false" max="1"/>
				<rule avp="Charging-Rule-Base-Name" required="false" max="1"/>
				<rule avp="Accounting-Input-Octets" required="false" max="1"/>
				<rule avp="Accounting-Output-Octets" required="false" max="1"/>
				<rule avp="Local-Sequence-Number" required="false" max="1"/>
				<rule avp="QoS-Information" required="false" max="1"/>
				<rule avp="Rating-Group" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
				<rule avp="Service-Identifier" required="false" max="1"/>
				<rule avp="Service-Specific-Info" required="false" max="1"/>
				<rule avp="ADC-Rule-Base-Name" required="false" max="1"/>
				<rule avp="SGSN-Address" required="false" max="1"/>
				<rule avp="Time-First-Usage" required="false" max="1"/>
				<rule avp="Time-Last-Usage" required="false" max="1"/>
				<rule avp="Time-Usage" required="false" max="1"/>
				<rule avp="Change-Condition" required="false"/>
				<rule avp="3GPP-User-Location-Info" required="false" max="1"/>
				<rule avp="3GPP2-BSID" required="false" max="1"/>
				<rule avp="UWAN-User-Location-Info" required="false" max="1"/>
				<rule avp="TWAN-User-Location-Info" required="false" max="1"/>
				<rule avp="Sponsor-Identity" required="false" max="1"/>
				<rule avp="Application-Service-Provider-Identity" required="false" max="1"/>
				<rule avp="Presence-Reporting-Area-Information" required="false"/>
				<rule avp="Presence-Reporting-Area-Status" required="false" max="1"/>
				<rule avp="User-CSG-Information" required="false" max="1"/>
				<rule avp="3GPP-RAT-Type" required="false" max="1"/>
				<rule avp="Related-Change-Condition-Information" required="false" max="1"/>
				<rule avp="Serving-PLMN-Rate-Control" required="false" max="1"/>
				<rule avp="APN-Rate-Control" required="false" max="1"/>
				<rule avp="3GPP-PS-Data-Off-Status" required="false" max="1"/>
				<rule avp="Traffic-Steering-Policy-Identifier-DL" required="false" max="1"/>
				<rule avp="Traffic-Steering-Policy-Identifier-UL" required="false" max="1"/>
				<rule avp="VoLTE-Information" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Terminal-Information" code="1401" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Grouped">
				<rule avp="IMEI" required="false" max="1"/>
				<rule avp="Software-Version" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Start-Time" code="2041" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Stop-Time" code="2042" must="V,M" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Change-Condition" code="2037" must="V,M" may="-" vendor-id="10415">
			<data type="Integer32"/>
		</avp>

		<avp name="Diagnostics" code="2039" must="V,M" may="-" vendor-id="10415">
			<data type="Integer32"/>
		</avp>

		<avp name="Low-Priority-Indicator" code="2602" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NO"/>
				<item code="1" name="YES"/>
			</data>
		</avp>

		<avp name="NBIFOM-Mode" code="2830" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="UE_INITIATED"/>
				<item code="1" name="NETWORK_INITIATED"/>
			</data>
		</avp>

		<avp name="NBIFOM-Support" code="2831" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NBIFOM_NOT_SUPPORTED"/>
				<item code="1" name="NBIFOM_SUPPORTED"/>
			</data>
		</avp>

		<avp name="MME-Number-for-MT-SMS" code="1645" must="V" must-not="M" may-encrypt="No" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="MME-Name" code="2402" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="MME-Realm" code="2408" must="V" must-not="M" may-encrypt="No" vendor-id="10415">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Logical-Access-ID" code="302" must="V" may="M" may-encrypt="Yes" vendor-id="13019">
			<data type="OctetString"/>
		</avp>

		<avp name="Physical-Access-ID" code="313" must="V" may="M" may-encrypt="Yes" vendor-id="13019">
			<data type="UTF8String"/>
		</avp>

		<avp name="Fixed-User-Location-Info" code="2825" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SSID" required="false" max="1"/>
				<rule avp="BSSID" required="false" max="1"/>
				<rule avp="Logical-Access-ID" required="false" max="1"/>
				<rule avp="Physical-Access-ID" required="false" max="1"/>
			</data>
		</avp>

		<avp name="CN-Operator-Selection-Entity" code="3421" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="THE_SERVING_NETWORK_HAS_BEEN_SELECTED_BY_THE_UE"/>
				<item code="1" name="THE_SERVING_NETWORK_HAS_BEEN_SELECTED_BY_THE_NETWORK"/>
			</data>
		</avp>

		<avp name="Enhanced-Diagnostics" code="3901" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="RAN-NAS-Release-Cause" required="false"/>
			</data>
		</avp>

		<avp name="SGi-PtP-Tunnelling-Method" code="3931" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="UDP_IP_BASED"/>
				<item code="1" name="OTHERS"/>
			</data>
		</avp>

		<avp name="CP-CIoT-EPS-Optimisation-Indicator" code="3930" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NOT_APPLY"/>
				<item code="1" name="APPLY"/>
			</data>
		</avp>

		<avp name="UNI-PDU-CP-Only-Flag" code="3932" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="UNI_PDU_BOTH_UP_CP"/>
				<item code="1" name="UNI_PDU_CP_ONLY"/>
			</data>
		</avp>

		<avp name="Serving-PLMN-Rate-Control" code="4310" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Uplink-Rate-Limit" required="false" max="1"/>
				<rule avp="Downlink-Rate-Limit" required="false" max="1"/>
			</data>
		</avp>

		<avp name="APN-Rate-Control" code="3933" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="APN-Rate-Control-Uplink" required="false" max="1"/>
				<rule avp="APN-Rate-Control-Downlink" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Charging-Per-IP-CAN-Session-Indicator" code="4400" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="INACTIVE"/>
				<item code="1" name="ACTIVE"/>
			</data>
		</avp>

		<avp name="RRC-Cause-Counter" code="4318" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Counter-Value" required="false" max="1"/>
				<rule avp="RRC-Counter-Timestamp" required="false" max="1"/>
			</data>
		</avp>

		<avp name="3GPP-PS-Data-Off-Status" code="4406" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="ACTIVE"/>
				<item code="1" name="INACTIVE"/>
			</data>
		</avp>

		<avp name="SCS-AS-Address" code="3940" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SCS-Realm" required="false" max="1"/>
				<rule avp="SCS-Address" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Unused-Quota-Timer" code="4407" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="RAN-Secondary-RAT-Usage-Report" code="1302" must="V" may="-" must-not="M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Secondary-RAT-Type" required="false" max="1"/>
				<rule avp="RAN-Start-Timestamp" required="false" max="1"/>
				<rule avp="RAN-End-Timestamp" required="false" max="1"/>
				<rule avp="Accounting-Input-Octets" required="false" max="1"/>
				<rule avp="Accounting-Output-Octets" required="false" max="1"/>
				<rule avp="3GPP-Charging-Id" required="false" max="1"/>
			</data>
		</avp>

		<avp name="CSG-Id" code="1437" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="CSG-Access-Mode" code="2317" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="CLOSED_MODE"/>
				<item code="1" name="HYBRID_MODE"/>
			</data>
		</avp>

		<avp name="CSG-Membership-Indication" code="2318" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NOT_CSG_MEMBER"/>
				<item code="1" name="CSG_MEMBER"/>
			</data>
		</avp>

		<avp name="SSID" code="1524" must="V" may="P" must-not="M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="BSSID" code="2716" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Civic-Address-Information" code="1305" must="V" may="-" must-not="M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="WLAN-Operator-Id" code="1306" must="V" may="-" must-not="M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="WLAN-PLMN-Id" required="false" max="1"/>
				<rule avp="WLAN-Operator-Name" required="false" max="1"/>
			</data>
		</avp>

		<avp name="WLAN-PLMN-Id" code="1308" must="V" may="-" must-not="M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="WLAN-Operator-Name" code="1307" must="V" may="-" must-not="M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Accounting-Input-Octets" code="363" must="M" may="-" must-not="V" vendor-id="0">
			<data type="Unsigned64"/>
		</avp>

		<avp name="Accounting-Output-Octets" code="364" must="M" may="-" must-not="V" vendor-id="0">
			<data type="Unsigned64"/>
		</avp>

		<avp name="Change-Time" code="2038" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Presence-Reporting-Area-Status" code="2823" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Access-Availability-Change-Reason" code="2833" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Related-Change-Condition-Information" code="3925" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SGSN-Address" required="false" max="1"/>
				<rule avp="Change-Condition" required="false"/>
				<rule avp="3GPP-User-Location-Info" required="false" max="1"/>
				<rule avp="3GPP2-BSID" required="false" max="1"/>
				<rule avp="UWAN-User-Location-Info" required="false" max="1"/>
				<rule avp="Presence-Reporting-Area-Status" required="false" max="1"/>
				<rule avp="User-CSG-Information" required="false" max="1"/>
				<rule avp="3GPP-RAT-Type" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Uplink-Rate-Limit" code="4311" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Downlink-Rate-Limit" code="4312" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="APN-Rate-Control-Uplink" code="3935" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Additional-Exception-Reports" required="false" max="1"/>
				<rule avp="Rate-Control-Time-Unit" required="false" max="1"/>
				<rule avp="Rate-Control-Max-Rate" required="false" max="1"/>
			</data>
		</avp>

		<avp name="APN-Rate-Control-Downlink" code="3934" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Rate-Control-Time-Unit" required="false" max="1"/>
				<rule avp="Rate-Control-Max-Rate" required="false" max="1"/>
				<rule avp="Rate-Control-Max-Message-Size" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Additional-Exception-Reports" code="3936" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NOT_ALLOWED"/>
				<item code="1" name="ALLOWED"/>
			</data>
		</avp>

		<avp name="Rate-Control-Time-Unit" code="3939" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Rate-Control-Max-Rate" code="3938" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Rate-Control-Max-Message-Size" code="3937" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Local-Sequence-Number" code="2063" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Time-First-Usage" code="2043" must="V,M" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Time-Last-Usage" code="2044" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Time-Usage" code="2045" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Sponsor-Identity" code="531" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Application-Service-Provider-Identity" code="532" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Traffic-Steering-Policy-Identifier-DL" code="2836" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Traffic-Steering-Policy-Identifier-UL" code="2837" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="VoLTE-Information" code="1323" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Calling-Party-Address" required="false" max="1"/>
				<rule avp="Callee-Information" required="false" max="1"/>
			</data>
		</avp>

		<avp name="IMEI" code="1402" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Software-Version" code="1403" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Counter-Value" code="4319" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="RRC-Counter-Timestamp" code="4320" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="SCS-Realm" code="3942" must="V,M" may="-" vendor-id="10415">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="SCS-Address" code="3941" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Secondary-RAT-Type" code="1304" must="V" may="-" must-not="M" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="RAN-Start-Timestamp" code="1303" must="V" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="RAN-End-Timestamp" code="1301" must="V" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Event-Type" code="823" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SIP-Method" required="false" max="1"/>
				<rule avp="Event" required="false" max="1"/>
				<rule avp="Expires" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Role-Of-Node" code="829" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="ORIGINATING_ROLE"/>
				<item code="1" name="TERMINATING_ROLE"/>
				<item code="2" name="FORWARDING_ROLE"/>
			</data>
		</avp>

		<avp name="Node-Functionality" code="862" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="S_CSCF"/>
				<item code="1" name="P_CSCF"/>
				<item code="2" name="I_CSCF"/>
				<item code="3" name="MRFC"/>
				<item code="4" name="MGCF"/>
				<item code="5" name="BGCF"/>
				<item code="6" name="AS"/>
				<item code="7" name="IBCF"/>
				<item code="8" name="S_GW"/>
				<item code="9" name="P_GW"/>
				<item code="10" name="HSGW"/>
				<item code="11" name="E_CSCF"/>
				<item code="12" name="MME"/>
				<item code="13" name="TRF"/>
				<item code="14" name="TF"/>
				<item code="15" name="ATCF"/>
				<item code="16" name="PROXY_FUNCTION"/>
				<item code="17" name="EPDG"/>
				<item code="18" name="TDF"/>
				<item code="19" name="TWAG"/>
				<item code="20" name="SCEF"/>
				<item code="21" name="IWK_SCEF"/>
			</data>
		</avp>

		<avp name="User-Session-Id" code="830" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Outgoing-Session-Id" code="2320" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Session-Priority" code="650" must="V" must-not="M" may-encrypt="No" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="PRIORITY_0"/>
				<item code="1" name="PRIORITY_1"/>
				<item code="2" name="PRIORITY_2"/>
				<item code="3" name="PRIORITY_3"/>
				<item code="4" name="PRIORITY_4"/>
			</data>
		</avp>

		<avp name="Called-Party-Address" code="832" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Called-Asserted-Identity" code="1250" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Called-Identity-Change" code="3917" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Called-Identity" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Number-Portability-Routing-Information" code="2024" must="V,M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Carrier-Select-Routing-Information" code="2023" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Alternate-Charged-Party-Address" code="1280" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Requested-Party-Address" code="1251" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Associated-URI" code="856" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Time-Stamps" code="833" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SIP-Request-Timestamp" required="false" max="1"/>
				<rule avp="SIP-Response-Timestamp" required="false" max="1"/>
				<rule avp="SIP-Request-Timestamp-Fraction" required="false" max="1"/>
				<rule avp="SIP-Response-Timestamp-Fraction" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Application-Server-Information" code="850" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Application-Server" required="false" max="1"/>
				<rule avp="Application-Provided-Called-Party-Address" required="false"/>
				<rule avp="Status-AS-Code" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Inter-Operator-Identifier" code="838" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Originating-IOI" required="false" max="1"/>
				<rule avp="Terminating-IOI" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Transit-IOI-List" code="2701" must="V,M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="IMS-Charging-Identifier" code="841" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="SDP-Session-Description" code="842" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="SDP-Media-Component" code="843" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SDP-Media-Name" required="false" max="1"/>
				<rule avp="SDP-Media-Description" required="false"/>
				<rule avp="Local-GW-Inserted-Indication" required="false" max="1"/>
				<rule avp="IP-Realm-Default-Indication" required="false" max="1"/>
				<rule avp="Transcoder-Inserted-Indication" required="false" max="1"/>
				<rule avp="Media-Initiator-Flag" required="false" max="1"/>
				<rule avp="Media-Initiator-Party" required="false" max="1"/>
				<rule avp="3GPP-Charging-Id" required="false" max="1"/>
				<rule avp="Access-Network-Charging-Identifier-Value" required="false" max="1"/>
				<rule avp="SDP-Type" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Served-Party-IP-Address" code="848" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Server-Capabilities" code="603" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Mandatory-Capability" required="false"/>
				<rule avp="Optional-Capability" required="false"/>
				<rule avp="Server-Name" required="false"/>
			</data>
		</avp>

		<avp name="Trunk-Group-ID" code="851" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Incoming-Trunk-Group-ID" required="false" max="1"/>
				<rule avp="Outgoing-Trunk-Group-ID" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Bearer-Service" code="854" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Service-Id" code="855" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Message-Body" code="889" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Content-Type" required="true" max="1"/>
				<rule avp="Content-Length" required="true" max="1"/>
				<rule avp="Content-Disposition" required="false" max="1"/>
				<rule avp="Originator" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Cause-Code" code="861" must="V,M" may="-" vendor-id="10415">
			<data type="Integer32"/>
		</avp>

		<avp name="Reason-Header" code="3401" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Access-Network-Information" code="1263" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Cellular-Network-Information" code="3924" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Early-Media-Description" code="1272" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SDP-TimeStamps" required="false" max="1"/>
				<rule avp="SDP-Media-Component" required="false"/>
				<rule avp="SDP-Session-Description" required="false"/>
			</data>
		</avp>

		<avp name="IMS-Communication-Service-Identifier" code="1281" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="IMS-Application-Reference-Identifier" code="2601" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Online-Charging-Flag" code="2303" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="ECF_ADDRESS_NOT_PROVIDED"/>
				<item code="1" name="ECF_ADDRESS_PROVIDED"/>
			</data>
		</avp>

		<avp name="Real-Time-Tariff-Information" code="2305" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Tariff-Information" required="false" max="1"/>
				<rule avp="Tariff-XML" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Account-Expiration" code="2309" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Initial-IMS-Charging-Identifier" code="2321" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="NNI-Information" code="2703" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Session-Direction" required="false" max="1"/>
				<rule avp="NNI-Type" required="false" max="1"/>
				<rule avp="Relationship-Mode" required="false" max="1"/>
				<rule avp="Neighbour-Node-Address" required="false" max="1"/>
			</data>
		</avp>

		<avp name="From-Address" code="2708" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="IMS-Emergency-Indicator" code="2322" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NON_EMERGENCY"/>
				<item code="1" name="EMERGENCY"/>
			</data>
		</avp>

		<avp name="IMS-Visited-Network-Identifier" code="2713" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Access-Network-Info-Change" code="4401" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Access-Network-Information" required="false"/>
				<rule avp="Cellular-Network-Information" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Access-Transfer-Information" code="2709" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Access-Transfer-Type" required="false" max="1"/>
				<rule avp="Access-Network-Information" required="false"/>
				<rule avp="Cellular-Network-Information" required="false" max="1"/>
				<rule avp="Inter-UE-Transfer" required="false" max="1"/>
				<rule avp="User-Equipment-Info" required="false" max="1"/>
				<rule avp="Instance-Id" required="false" max="1"/>
				<rule avp="Related-IMS-Charging-Identifier" required="false" max="1"/>
				<rule avp="Related-IMS-Charging-Identifier-Node" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Related-IMS-Charging-Identifier" code="2711" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Related-IMS-Charging-Identifier-Node" code="2712" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Route-Header-Received" code="3403" must="V,M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Route-Header-Transmitted" code="3404" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Instance-Id" code="3402" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="TAD-Identifier" code="2717" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="CS"/>
				<item code="1" name="PS"/>
			</data>
		</avp>

		<avp name="FE-Identifier-List" code="4413" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="SIP-Method" code="824" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Event" code="825" must="V,M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Expires" code="888" must="V,M" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Called-Identity" code="3916" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="SIP-Request-Timestamp" code="834" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="SIP-Response-Timestamp" code="835" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="SIP-Request-Timestamp-Fraction" code="2301" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="SIP-Response-Timestamp-Fraction" code="2302" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Application-Server" code="836" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Application-Provided-Called-Party-Address" code="837" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Status-AS-Code" code="2702" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="4XX"/>
				<item code="1" name="5XX"/>
				<item code="2" name="TIMEOUT"/>
			</data>
		</avp>

		<avp name="Originating-IOI" code="839" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Terminating-IOI" code="840" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="SDP-Media-Name" code="844" must="V,M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="SDP-Media-Description" code="845" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Local-GW-Inserted-Indication" code="2604" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="LOCAL_GW_NOT_INSERTED"/>
				<item code="1" name="LOCAL_GW_INSERTED"/>
			</data>
		</avp>

		<avp name="IP-Realm-Default-Indication" code="2603" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="DEFAULT_IP_REALM_NOT_USED"/>
				<item code="1" name="DEFAULT_IP_REALM_USED"/>
			</data>
		</avp>

		<avp name="Transcoder-Inserted-Indication" code="2605" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="TRANSCODER_NOT_INSERTED"/>
				<item code="1" name="TRANSCODER_INSERTED"/>
			</data>
		</avp>

		<avp name="Media-Initiator-Flag" code="882" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="CALLED_PARTY"/>
				<item code="1" name="CALLING_PARTY"/>
				<item code="2" name="UNKNOWN"/>
			</data>
		</avp>

		<avp name="Media-Initiator-Party" code="1288" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Access-Network-Charging-Identifier-Value" code="503" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="SDP-Type" code="2036" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="SDP_OFFER"/>
				<item code="1" name="SDP_ANSWER"/>
			</data>
		</avp>

		<avp name="Mandatory-Capability" code="604" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Optional-Capability" code="605" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Server-Name" code="602" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Incoming-Trunk-Group-ID" code="852" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Outgoing-Trunk-Group-ID" code="853" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Content-Type" code="826" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Content-Length" code="827" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Content-Disposition" code="828" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Originator" code="864" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="CALLING_PARTY"/>
				<item code="1" name="CALLED_PARTY"/>
			</data>
		</avp>

		<avp name="SDP-TimeStamps" code="1273" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SDP-Offer-Timestamp" required="false" max="1"/>
				<rule avp="SDP-Answer-Timestamp" required="false" max="1"/>
			</data>
		</avp>

		<avp name="SDP-Offer-Timestamp" code="1274" must="V,M" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="SDP-Answer-Timestamp" code="1275" must="V,M" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Tariff-XML" code="2306" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Session-Direction" code="2707" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="INBOUND"/>
				<item code="1" name="OUTBOUND"/>
			</data>
		</avp>

		<avp name="NNI-Type" code="2704" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NON_ROAMING"/>
				<item code="1" name="ROAMING_WITHOUT_LOOPBACK"/>
				<item code="2" name="ROAMING_WITH_LOOPBACK"/>
			</data>
		</avp>

		<avp name="Relationship-Mode" code="2706" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="TRUSTED"/>
				<item code="1" name="NON_TRUSTED"/>
			</data>
		</avp>

		<avp name="Neighbour-Node-Address" code="2705" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Access-Transfer-Type" code="2710" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="PS_TO_CS_TRANSFER"/>
				<item code="1" name="CS_TO_PS_TRANSFER"/>
				<item code="2" name="PS_TO_PS_TRANSFER"/>
				<item code="3" name="CS_TO_CS_TRANSFER"/>
			</data>
		</avp>

		<avp name="Inter-UE-Transfer" code="3902" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="INTRA_UE_TRANSFER"/>
				<item code="1" name="INTER_UE_TRANSFER"/>
			</data>
		</avp>

		<avp name="Originator-Address" code="886" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Address-Type" required="false" max="1"/>
				<rule avp="Address-Data" required="false" max="1"/>
				<rule avp="Address-Domain" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Recipient-Address" code="1201" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Address-Type" required="false" max="1"/>
				<rule avp="Address-Data" required="false" max="1"/>
				<rule avp="Address-Domain" required="false" max="1"/>
				<rule avp="Addressee-Type" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Submission-Time" code="1202" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="MM-Content-Type" code="1203" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Type-Number" required="false" max="1"/>
				<rule avp="Additional-Type-Information" required="false" max="1"/>
				<rule avp="Content-Size" required="false" max="1"/>
				<rule avp="Additional-Content-Information" required="false"/>
			</data>
		</avp>

		<avp name="Priority" code="1209" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="LOW"/>
				<item code="1" name="NORMAL"/>
				<item code="2" name="HIGH"/>
			</data>
		</avp>

		<avp name="Message-ID" code="1210" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Message-Type" code="1211" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="1" name="M_SEND_REQ"/>
				<item code="2" name="M_SEND_CONF"/>
				<item code="3" name="M_NOTIFICATION_IND"/>
				<item code="4" name="M_NOTIFYRESP_IND"/>
				<item code="5" name="M_RETRIEVE_CONF"/>
				<item code="6" name="M_ACKNOWLEDGE_IND"/>
				<item code="7" name="M_DELIVERY_IND"/>
				<item code="8" name="M_READ_REC_IND"/>
				<item code="9" name="M_READ_ORIG_IND"/>
				<item code="10" name="M_FORWARD_REQ"/>
				<item code="11" name="M_FORWARD_CONF"/>
				<item code="12" name="M_MBOX_STORE_CONF"/>
				<item code="13" name="M_MBOX_VIEW_CONF"/>
				<item code="14" name="M_MBOX_UPLOAD_CONF"/>
				<item code="15" name="M_MBOX_DELETE_CONF"/>
			</data>
		</avp>

		<avp name="Message-Size" code="1212" must="V,M" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Message-Class" code="1213" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Class-Identifier" required="false" max="1"/>
				<rule avp="Token-Text" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Delivery-Report-Requested" code="1216" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NO"/>
				<item code="1" name="YES"/>
			</data>
		</avp>

		<avp name="Read-Reply-Report-Requested" code="1222" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NO"/>
				<item code="1" name="YES"/>
			</data>
		</avp>

		<avp name="MMBox-Storage-Requested" code="1248" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NO"/>
				<item code="1" name="YES"/>
			</data>
		</avp>

		<avp name="Applic-ID" code="1218" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Reply-Applic-ID" code="1223" must="V,M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Aux-Applic-Info" code="1219" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Content-Class" code="1220" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="TEXT"/>
				<item code="1" name="IMAGE_BASIC"/>
				<item code="2" name="IMAGE_RICH"/>
				<item code="3" name="VIDEO_BASIC"/>
				<item code="4" name="VIDEO_RICH"/>
				<item code="5" name="MEGAPIXEL"/>
				<item code="6" name="CONTENT_BASIC"/>
				<item code="7" name="CONTENT_RICH"/>
			</data>
		</avp>

		<avp name="DRM-Content" code="1221" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NO"/>
				<item code="1" name="YES"/>
			</data>
		</avp>

		<avp name="Adaptations" code="1217" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="YES"/>
				<item code="1" name="NO"/>
			</data>
		</avp>

		<avp name="VASP-Id" code="1101" must="V,M" may="P" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="VAS-Id" code="1102" must="V,M" may="P" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Address-Type" code="899" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="E_MAIL_ADDRESS"/>
				<item code="1" name="MSISDN"/>
				<item code="2" name="IPV4_ADDRESS"/>
				<item code="3" name="IPV6_ADDRESS"/>
				<item code="4" name="NUMERIC_SHORTCODE"/>
				<item code="5" name="ALPHANUMERIC_SHORTCODE"/>
				<item code="6" name="OTHER"/>
				<item code="7" name="IMSI"/>
			</data>
		</avp>

		<avp name="Address-Data" code="897" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Address-Domain" code="898" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Domain-Name" required="false" max="1"/>
				<rule avp="3GPP-IMSI-MCC-MNC" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Domain-Name" code="1200" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Addressee-Type" code="1208" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="TO"/>
				<item code="1" name="CC"/>
				<item code="2" name="BCC"/>
			</data>
		</avp>

		<avp name="Additional-Type-Information" code="1205" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Content-Size" code="1206" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Additional-Content-Information" code="1207" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Type-Number" required="false" max="1"/>
				<rule avp="Additional-Type-Information" required="false" max="1"/>
				<rule avp="Content-Size" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Class-Identifier" code="1214" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="PERSONAL"/>
				<item code="1" name="ADVERTISEMENT"/>
				<item code="2" name="INFORMATIONAL"/>
				<item code="3" name="AUTO"/>
			</data>
		</avp>

		<avp name="Token-Text" code="1215" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="LCS-Client-ID" code="1232" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="LCS-Client-Type" required="false" max="1"/>
				<rule avp="LCS-Client-External-ID" required="false" max="1"/>
				<rule avp="LCS-Client-Dialed-By-MS" required="false" max="1"/>
				<rule avp="LCS-Client-Name" required="false" max="1"/>
				<rule avp="LCS-APN" required="false" max="1"/>
				<rule avp="LCS-Requestor-ID" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Location-Type" code="1244" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Location-Estimate-Type" required="false" max="1"/>
				<rule avp="Deferred-Location-Event-Type" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Location-Estimate" code="1242" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Positioning-Data" code="1245" must="V,M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="3GPP-IMSI" code="1" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="MSISDN" code="701" must="V,M" may="P" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="LCS-Client-Type" code="1241" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="EMERGENCY_SERVICES"/>
				<item code="1" name="VALUE_ADDED_SERVICES"/>
				<item code="2" name="PLMN_OPERATOR_SERVICES"/>
				<item code="3" name="LAWFUL_INTERCEPT_SERVICES"/>
			</data>
		</avp>

		<avp name="LCS-Client-External-ID" code="1234" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="LCS-Client-Dialed-By-MS" code="1233" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="LCS-Client-Name" code="1235" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="LCS-Data-Coding-Scheme" required="false" max="1"/>
				<rule avp="LCS-Name-String" required="false" max="1"/>
				<rule avp="LCS-Format-Indicator" required="false" max="1"/>
			</data>
		</avp>

		<avp name="LCS-APN" code="1231" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="LCS-Requestor-ID" code="1239" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="LCS-Data-Coding-Scheme" required="false" max="1"/>
				<rule avp="LCS-Requestor-ID-String" required="false" max="1"/>
			</data>
		</avp>

		<avp name="LCS-Data-Coding-Scheme" code="1236" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="LCS-Name-String" code="1238" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="LCS-Format-Indicator" code="1237" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="LOGICAL_NAME"/>
				<item code="1" name="EMAIL_ADDRESS"/>
				<item code="2" name="MSISDN"/>
				<item code="3" name="URL"/>
				<item code="4" name="SIP_URL"/>
			</data>
		</avp>

		<avp name="LCS-Requestor-ID-String" code="1240" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Location-Estimate-Type" code="1243" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="CURRENT_LOCATION"/>
				<item code="1" name="CURRENT_LAST_KNOWN_LOCATION"/>
				<item code="2" name="INITIAL_LOCATION"/>
				<item code="3" name="ACTIVATE_DEFERRED_LOCATION"/>
				<item code="4" name="CANCEL_DEFERRED_LOCATION"/>
			</data>
		</avp>

		<avp name="Deferred-Location-Event-Type" code="1230" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="PoC-Server-Role" code="883" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="PARTICIPATING_POC_SERVER"/>
				<item code="1" name="CONTROLLING_POC_SERVER"/>
				<item code="2" name="INTERWORKING_FUNCTION"/>
				<item code="3" name="INTERWORKING_SELECTION_FUNCTION"/>
			</data>
		</avp>

		<avp name="PoC-Session-Type" code="884" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="1_TO_1_POC_SESSION"/>
				<item code="1" name="CHAT_POC_GROUP_SESSION"/>
				<item code="2" name="PRE_ARRANGED_POC_GROUP_SESSION"/>
				<item code="3" name="AD_HOC_POC_GROUP_SESSION"/>
			</data>
		</avp>

		<avp name="PoC-User-Role" code="1252" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="PoC-User-Role-IDs" required="false" max="1"/>
				<rule avp="PoC-User-Role-Info-Units" required="false" max="1"/>
			</data>
		</avp>

		<avp name="PoC-Session-Initiation-Type" code="1277" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="PRE_ESTABLISHED"/>
				<item code="1" name="ON_DEMAND"/>
			</data>
		</avp>

		<avp name="PoC-Event-Type" code="2025" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NORMAL"/>
				<item code="1" name="INSTANT_PPERSONAL_AALERT_EVENT"/>
				<item code="2" name="POC_GROUP_ADVERTISEMENT_EVENT"/>
				<item code="3" name="EARLY_SSESSION_SETTING_UP_EVENT"/>
				<item code="4" name="POC_TALK_BURST"/>
			</data>
		</avp>

		<avp name="Number-Of-Participants" code="885" must="V,M" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Participants-Involved" code="887" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Participant-Group" code="1260" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Called-Party-Address" required="false" max="1"/>
				<rule avp="Participant-Access-Priority" required="false" max="1"/>
				<rule avp="User-Participating-Type" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Talk-Burst-Exchange" code="1255" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="PoC-Change-Time" required="true" max="1"/>
				<rule avp="Number-Of-Talk-Bursts" required="false" max="1"/>
				<rule avp="Talk-Burst-Volume" required="false" max="1"/>
				<rule avp="Talk-Burst-Time" required="false" max="1"/>
				<rule avp="Number-Of-Received-Talk-Bursts" required="false" max="1"/>
				<rule avp="Received-Talk-Burst-Volume" required="false" max="1"/>
				<rule avp="Received-Talk-Burst-Time" required="false" max="1"/>
				<rule avp="Number-Of-Participants" required="false" max="1"/>
				<rule avp="PoC-Change-Condition" required="false" max="1"/>
			</data>
		</avp>

		<avp name="PoC-Controlling-Address" code="858" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="PoC-Group-Name" code="859" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="PoC-Session-Id" code="1229" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Charged-Party" code="857" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="PoC-User-Role-IDs" code="1253" must="V,M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="PoC-User-Role-Info-Units" code="1254" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="1" name="MODERATOR"/>
				<item code="2" name="DISPATCHER"/>
				<item code="3" name="SESSION_OWNER"/>
				<item code="4" name="SESSION_PARTICIPANT"/>
			</data>
		</avp>

		<avp name="Participant-Access-Priority" code="1259" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="1" name="PRE_EMPTIVE_PRIORITY"/>
				<item code="2" name="HIGH_PRIORITY"/>
				<item code="3" name="NORMAL_PRIORITY"/>
				<item code="4" name="LOW_PRIORITY"/>
			</data>
		</avp>

		<avp name="User-Participating-Type" code="1279" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NORMAL"/>
				<item code="1" name="NW_POC_BOX"/>
				<item code="2" name="UE_POC_BOX"/>
			</data>
		</avp>

		<avp name="PoC-Change-Time" code="1262" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Number-Of-Talk-Bursts" code="1283" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Talk-Burst-Volume" code="1287" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Talk-Burst-Time" code="1286" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Number-Of-Received-Talk-Bursts" code="1282" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Received-Talk-Burst-Volume" code="1285" must="V,M" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Received-Talk-Burst-Time" code="1284" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="PoC-Change-Condition" code="1261" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="SERVICECHANGE"/>
				<item code="1" name="VOLUMELIMIT"/>
				<item code="2" name="TIMELIMIT"/>
				<item code="3" name="NUMBEROFTALKBURSTLIMIT"/>
				<item code="4" name="NUMBEROFACTIVEPARTICIPANTS"/>
				<item code="5" name="TARIFFTIME"/>
			</data>
		</avp>

		<avp name="TMGI" code="900" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="MBMS-Service-Type" code="906" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="MULTICAST"/>
				<item code="1" name="BROADCAST"/>
			</data>
		</avp>

		<avp name="MBMS-User-Service-Type" code="1225" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="1" name="DOWNLOAD"/>
				<item code="2" name="STREAMING"/>
			</data>
		</avp>

		<avp name="File-Repair-Supported" code="1224" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="1" name="SUPPORTED"/>
				<item code="2" name="NOT_SUPPORTED"/>
			</data>
		</avp>

		<avp name="Required-MBMS-Bearer-Capabilities" code="901" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="MBMS-2G-3G-Indicator" code="907" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="2G"/>
				<item code="1" name="3G"/>
				<item code="2" name="2G_AND_3G"/>
			</data>
		</avp>

		<avp name="RAI" code="909" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="MBMS-Service-Area" code="903" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="MBMS-Session-Identity" code="908" must="V,M" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="CN-IP-Multicast-Distribution" code="921" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NO_IP_MULTICAST"/>
				<item code="1" name="IP_MULTICAST"/>
			</data>
		</avp>

		<avp name="MBMS-GW-Address" code="2307" must="V,M" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="MBMS-Charged-Party" code="2323" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="CONTENT_PROVIDER"/>
				<item code="1" name="SUBSCRIBER"/>
			</data>
		</avp>

		<avp name="MBMS-Data-Transfer-Start" code="929" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned64"/>
		</avp>

		<avp name="MBMS-Data-Transfer-Stop" code="930" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned64"/>
		</avp>

		<avp name="SMS-Node" code="2016" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="SMS_ROUTER"/>
				<item code="1" name="IP_SM_GW"/>
				<item code="2" name="SMS_ROUTER_AND_IP_SM_GW"/>
				<item code="3" name="SMS_SC"/>
			</data>
		</avp>

		<avp name="Client-Address" code="2018" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Originator-SCCP-Address" code="2008" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="SMSC-Address" code="2017" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Data-Coding-Scheme" code="2001" must="V,M" may="-" vendor-id="10415">
			<data type="Integer32"/>
		</avp>

		<avp name="SM-Discharge-Time" code="2012" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="SM-Message-Type" code="2007" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="SUBMISSION"/>
				<item code="1" name="DELIVERY_REPORT"/>
				<item code="2" name="SM_SERVICE_REQUEST"/>
				<item code="3" name="T4_DEVICE_TRIGGER"/>
				<item code="4" name="SM_DEVICE_TRIGGER"/>
				<item code="5" name="MO_SMS_T4_SUBMISSION"/>
			</data>
		</avp>

		<avp name="Originator-Interface" code="2009" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Interface-Id" required="false" max="1"/>
				<rule avp="Interface-Text" required="false" max="1"/>
				<rule avp="Interface-Port" required="false" max="1"/>
				<rule avp="Interface-Type" required="false" max="1"/>
			</data>
		</avp>

		<avp name="SM-Protocol-ID" code="2013" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Reply-Path-Requested" code="2011" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NO_REPLY_PATH_SET"/>
				<item code="1" name="REPLY_PATH_SET"/>
			</data>
		</avp>

		<avp name="SM-Status" code="2014" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="SM-User-Data-Header" code="2015" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Number-Of-Messages-Sent" code="2019" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="SM-Sequence-Number" code="3408" must="V,M" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Recipient-Info" code="2026" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Destination-Interface" required="false" max="1"/>
				<rule avp="Recipient-Address" required="false"/>
				<rule avp="Recipient-Received-Address" required="false"/>
				<rule avp="Recipient-SCCP-Address" required="false" max="1"/>
				<rule avp="SM-Protocol-ID" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Originator-Received-Address" code="2027" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Address-Type" required="false" max="1"/>
				<rule avp="Address-Data" required="false" max="1"/>
				<rule avp="Address-Domain" required="false" max="1"/>
			</data>
		</avp>

		<avp name="SM-Service-Type" code="2029" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="VAS4SMS_SHORT_MESSAGE_CONTENT_PROCESSING"/>
				<item code="1" name="VAS4SMS_SHORT_MESSAGE_FORWARDING"/>
				<item code="2" name="VAS4SMS_SHORT_MESSAGE_FORWARDING_MULTIPLE_SUBSCRIPTIONS"/>
				<item code="3" name="VAS4SMS_SHORT_MESSAGE_FILTERING"/>
				<item code="4" name="VAS4SMS_SHORT_MESSAGE_RECEIPT"/>
				<item code="5" name="VAS4SMS_SHORT_MESSAGE_NETWORK_STORAGE"/>
				<item code="6" name="VAS4SMS_SHORT_MESSAGE_TO_MULTIPLE_DESTINATIONS"/>
				<item code="7" name="VAS4SMS_SHORT_MESSAGE_VIRTUAL_PRIVATE_NETWORK_VPN"/>
				<item code="8" name="VAS4SMS_SHORT_MESSAGE_AUTO_REPLY"/>
				<item code="9" name="VAS4SMS_SHORT_MESSAGE_PERSONAL_SIGNATURE"/>
				<item code="10" name="VAS4SMS_SHORT_MESSAGE_DEFERRED_DELIVERY"/>
			</data>
		</avp>

		<avp name="SMS-Result" code="3409" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="SM-Device-Trigger-Indicator" code="3407" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NOT_DEVICETRIGGER"/>
				<item code="1" name="DEVICE_TRIGGER_REQUEST"/>
				<item code="2" name="DEVICE_TRIGGER_REPLACE"/>
				<item code="3" name="DEVICE_TRIGGER_RECALL"/>
			</data>
		</avp>

		<avp name="SM-Device-Trigger-Information" code="3405" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="MTC-IWF-Address" required="false" max="1"/>
				<rule avp="Reference-Number" required="false" max="1"/>
				<rule avp="Serving-Node" required="false" max="1"/>
				<rule avp="Validity-Time" required="false" max="1"/>
				<rule avp="Priority-Indication" required="false" max="1"/>
				<rule avp="Application-Port-Identifier" required="false" max="1"/>
			</data>
		</avp>

		<avp name="MTC-IWF-Address" code="3406" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Application-Port-Identifier" code="3010" must="M,V" may="P" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="External-Identifier" code="3111" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Interface-Id" code="2003" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Interface-Text" code="2005" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Interface-Port" code="2004" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Interface-Type" code="2006" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="UNKNOWN"/>
				<item code="1" name="MOBILE_ORIGINATING"/>
				<item code="2" name="MOBILE_TERMINATING"/>
				<item code="3" name="APPLICATION_ORIGINATING"/>
				<item code="4" name="APPLICATION_TERMINATION"/>
			</data>
		</avp>

		<avp name="Destination-Interface" code="2002" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Interface-Id" required="false" max="1"/>
				<rule avp="Interface-Text" required="false" max="1"/>
				<rule avp="Interface-Port" required="false" max="1"/>
				<rule avp="Interface-Type" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Recipient-Received-Address" code="2028" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Address-Type" required="false" max="1"/>
				<rule avp="Address-Data" required="false" max="1"/>
				<rule avp="Address-Domain" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Recipient-SCCP-Address" code="2010" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Reference-Number" code="3007" must="M,V" may="P" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Serving-Node" code="2401" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Grouped">
				<rule avp="SGSN-Number" required="false" max="1"/>
				<rule avp="SGSN-Name" required="false" max="1"/>
				<rule avp="SGSN-Realm" required="false" max="1"/>
				<rule avp="MME-Name" required="false" max="1"/>
				<rule avp="MME-Realm" required="false" max="1"/>
				<rule avp="MSC-Number" required="false" max="1"/>
				<rule avp="3GPP-AAA-Server-Name" required="false" max="1"/>
				<rule avp="LCS-Capabilities-Sets" required="false" max="1"/>
				<rule avp="GMLC-Address" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Priority-Indication" code="3006" must="M,V" may="P" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NON_PRIORITY"/>
				<item code="1" name="PRIORITY"/>
			</data>
		</avp>

		<avp name="SGSN-Number" code="1489" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="SGSN-Name" code="2409" must="V" must-not="M" may-encrypt="No" vendor-id="10415">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="SGSN-Realm" code="2410" must="V" must-not="M" may-encrypt="No" vendor-id="10415">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="MSC-Number" code="2403" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="3GPP-AAA-Server-Name" code="318" must="V,M" may="P" vendor-id="10415">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="LCS-Capabilities-Sets" code="2404" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="GMLC-Address" code="2405" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Bearer-Capability" code="3412" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Network-Call-Reference-Number" code="3418" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="MSC-Address" code="3417" must="V,M" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Basic-Service-Code" code="3411" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Bearer-Service" required="false" max="1"/>
				<rule avp="Teleservice" required="false" max="1"/>
			</data>
		</avp>

		<avp name="ISUP-Location-Number" code="3414" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="VLR-Number" code="3420" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Forwarding-Pending" code="3415" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="FORWARDING_NOT_PENDING"/>
				<item code="1" name="FORWARDING_PENDING"/>
			</data>
		</avp>

		<avp name="ISUP-Cause" code="3416" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="ISUP-Cause-Location" required="false" max="1"/>
				<rule avp="ISUP-Cause-Value" required="false" max="1"/>
				<rule avp="ISUP-Cause-Diagnostics" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Start-of-Charging" code="3419" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="PS-Free-Format-Data" code="866" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Teleservice" code="3413" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="ISUP-Cause-Location" code="3423" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="ISUP-Cause-Value" code="3424" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="ISUP-Cause-Diagnostics" code="3422" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Supplementary-Service" code="2048" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="MMTel-SService-Type" required="false" max="1"/>
				<rule avp="Service-Mode" required="false" max="1"/>
				<rule avp="Number-Of-Diversions" required="false" max="1"/>
				<rule avp="Associated-Party-Address" required="false" max="1"/>
				<rule avp="Service-ID" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
				<rule avp="Number-Of-Participants" required="false" max="1"/>
				<rule avp="Participant-Action-Type" required="false" max="1"/>
				<rule avp="CUG-Information" required="false" max="1"/>
				<rule avp="AoC-Information" required="false" max="1"/>
			</data>
		</avp>

		<avp name="MMTel-SService-Type" code="2031" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Service-Mode" code="2032" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Number-Of-Diversions" code="2034" must="V,M" may="-" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Associated-Party-Address" code="2035" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Participant-Action-Type" code="2049" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="CREATE_CONF"/>
				<item code="1" name="JOIN_CONF"/>
				<item code="2" name="INVITE_INTO_CONF"/>
				<item code="3" name="QUIT_CONF"/>
			</data>
		</avp>

		<avp name="CUG-Information" code="2304" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Announcing-UE-HPLMN-Identifier" code="3426" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Announcing-UE-VPLMN-Identifier" code="3427" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Monitoring-UE-HPLMN-Identifier" code="3431" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Monitoring-UE-VPLMN-Identifier" code="3433" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Monitored-PLMN-Identifier" code="3430" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Role-Of-ProSe-Function" code="3438" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="HPLMN"/>
				<item code="1" name="VPLMN"/>
				<item code="2" name="LOCAL_PLMN"/>
			</data>
		</avp>

		<avp name="ProSe-App-Id" code="3811" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="ProSe-3rd-Party-Application-ID" code="3440" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Application-Specific-Data" code="3458" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="ProSe-Event-Type" code="3443" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="ANNOUCING"/>
				<item code="1" name="MONITORING"/>
				<item code="2" name="MATCH_REPORT"/>
			</data>
		</avp>

		<avp name="ProSe-Direct-Discovery-Model" code="3442" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="MODEL_A"/>
				<item code="1" name="MODEL_B"/>
			</data>
		</avp>

		<avp name="ProSe-Function-IP-Address" code="3444" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="ProSe-Function-ID" code="3602" must="M,V" may="P" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="ProSe-Validity-Timer" code="3815" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="ProSe-Role-Of-UE" code="3451" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="ANNOUNCING_UE"/>
				<item code="1" name="MONITORING_UE"/>
				<item code="2" name="REQUESTOR_UE"/>
				<item code="3" name="REQUESTED_UE"/>
			</data>
		</avp>

		<avp name="ProSe-Request-Timestamp" code="3450" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="PC3-Control-Protocol-Cause" code="3434" must="V,M" may="-" vendor-id="10415">
			<data type="Integer32"/>
		</avp>

		<avp name="Monitoring-UE-Identifier" code="3432" must="V,M" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="ProSe-Function-PLMN-Identifier" code="3457" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Requestor-PLMN-Identifier" code="3437" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Origin-App-Layer-User-Id" code="3600" must="M,V" may="P" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="WLAN-Link-Layer-Id" code="3821" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Requesting-EPUID" code="3816" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Target-App-Layer-User-Id" code="3601" must="M,V" may="P" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Requested-PLMN-Identifier" code="3436" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Time-Window" code="3818" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="ProSe-Range-Class" code="3448" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="RESERVED"/>
				<item code="1" name="50_M"/>
				<item code="2" name="100_M"/>
				<item code="3" name="200_M"/>
				<item code="4" name="500_M"/>
				<item code="5" name="1000_M"/>
				<item code="255" name="UNUSED"/>
			</data>
		</avp>

		<avp name="Proximity-Alert-Indication" code="3454" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="ALERT"/>
				<item code="1" name="NO_ALERT"/>
			</data>
		</avp>

		<avp name="Proximity-Alert-Timestamp" code="3455" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Proximity-Cancellation-Timestamp" code="3456" must="V,M" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="ProSe-Reason-For-Cancellation" code="3449" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="PROXIMITY_ALERT_SENT"/>
				<item code="1" name="TIME_EXPIRED_WITH_NO_RENEWAL"/>
				<item code="2" name="REQUESTOR_CANCELLATION"/>
			</data>
		</avp>

		<avp name="PC3-EPC-Control-Protocol-Cause" code="3435" must="V,M" may="-" vendor-id="10415">
			<data type="Integer32"/>
		</avp>

		<avp name="ProSe-UE-ID" code="3453" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="ProSe-Source-IP-Address" code="3452" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Layer-2-Group-ID" code="3429" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="ProSe-Group-IP-Multicast-Address" code="3446" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="Coverage-Info" code="3459" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Coverage-Status" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
				<rule avp="Location-Info" required="false"/>
			</data>
		</avp>

		<avp name="Radio-Parameter-Set-Info" code="3463" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Radio-Parameter-Set-Values" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Transmitter-Info" code="3468" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="ProSe-Source-IP-Address" required="false" max="1"/>
				<rule avp="ProSe-UE-ID" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Time-First-Transmission" code="3467" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Time-First-Reception" code="3466" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="ProSe-Direct-Communication-Transmission-Data-Container" code="3441" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Local-Sequence-Number" required="false" max="1"/>
				<rule avp="Coverage-Status" required="false" max="1"/>
				<rule avp="3GPP-User-Location-Info" required="false" max="1"/>
				<rule avp="Accounting-Output-Octets" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
				<rule avp="Change-Condition" required="false" max="1"/>
				<rule avp="Visited-PLMN-Id" required="false" max="1"/>
				<rule avp="Usage-Information-Report-Sequence-Number" required="false" max="1"/>
				<rule avp="Radio-Resources-Indicator" required="false" max="1"/>
				<rule avp="Radio-Frequency" required="false" max="1"/>
			</data>
		</avp>

		<avp name="ProSe-Direct-Communication-Reception-Data-Container" code="3461" must="V,M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Local-Sequence-Number" required="false" max="1"/>
				<rule avp="Coverage-Status" required="false" max="1"/>
				<rule avp="3GPP-User-Location-Info" required="false" max="1"/>
				<rule avp="Accounting-Input-Octets" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
				<rule avp="Change-Condition" required="false" max="1"/>
				<rule avp="Visited-PLMN-Id" required="false" max="1"/>
				<rule avp="Usage-Information-Report-Sequence-Number" required="false" max="1"/>
				<rule avp="Radio-Resources-Indicator" required="false" max="1"/>
				<rule avp="Radio-Frequency" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Announcing-PLMN-ID" code="4408" must="V,M" may="-" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="ProSe-Target-Layer-2-ID" code="4410" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Relay-IP-address" code="4411" must="V,M" may="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="ProSe-UE-to-Network-Relay-UE-ID" code="4409" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Target-IP-Address" code="4412" must="VM" may="-" must-not="-" vendor-id="10415">
			<data type="Address"/>
		</avp>

		<avp name="PC5-Radio-Technology" code="1300" must="V" may="-" must-not="M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="EUTRA"/>
				<item code="1" name="WLAN"/>
				<item code="2" name="BOTH_EUTRA_AND_WLAN"/>
			</data>
		</avp>

		<avp name="Coverage-Status" code="3428" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="OUT_OF_COVERAGE"/>
				<item code="1" name="IN_COVERAGE"/>
			</data>
		</avp>

		<avp name="Location-Info" code="3460" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="3GPP-User-Location-Info" required="false" max="1"/>
				<rule avp="Change-Time" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Radio-Parameter-Set-Values" code="3464" must="V,M" may="-" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Visited-PLMN-Id" code="1407" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Usage-Information-Report-Sequence-Number" code="3439" must="V,M" may="-" vendor-id="10415">
			<data type="Integer32"/>
		</avp>

		<avp name="Radio-Resources-Indicator" code="3465" must="V,M" may="-" vendor-id="10415">
			<data type="Integer32"/>
		</avp>

		<avp name="Radio-Frequency" code="3462" must="V,M" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Application-Server-ID" code="2101" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Application-Session-ID" code="2103" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Delivery-Status" code="2104" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Total-Number-Of-Messages-Sent" code="2114" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Total-Number-Of-Messages-Exploded" code="2113" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Number-Of-Messages-Successfully-Sent" code="2112" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Number-Of-Messages-Successfully-Exploded" code="2111" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Content-ID" code="2116" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Content-Provider-ID" code="2117" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="UTF8String"/>
		</avp>

		<avp name="Application-Entity-ID" code="1001" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="UTF8String"/>
		</avp>

		<avp name="External-ID" code="1005" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="UTF8String"/>
		</avp>

		<avp name="Receiver" code="1014" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="UTF8String"/>
		</avp>

		<avp name="Hosting-CSE-ID" code="1007" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="UTF8String"/>
		</avp>

		<avp name="Target-ID" code="1022" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="UTF8String"/>
		</avp>

		<avp name="Protocol-Type" code="1013" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Enumerated">
				<item code="0" name="HTTP"/>
				<item code="1" name="COAP"/>
				<item code="2" name="MQTT"/>
			</data>
		</avp>

		<avp name="Request-Operation" code="1017" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Enumerated">
				<item code="1" name="CREATE"/>
				<item code="2" name="RETRIEVE"/>
				<item code="3" name="UPDATE"/>
				<item code="4" name="DELETE"/>
				<item code="5" name="NOTIFY"/>
			</data>
		</avp>

		<avp name="Request-Headers-Size" code="1016" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Request-Body-Size" code="1015" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Response-Headers-Size" code="1019" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Response-Body-Size" code="1018" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Response-Status-Code" code="1020" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Enumerated">
				<item code="1000" name="ACCEPTED"/>
				<item code="2000" name="OK"/>
				<item code="2001" name="CREATED"/>
				<item code="2002" name="DELETED"/>
				<item code="2004" name="UPDATED"/>
				<item code="4000" name="BAD_REQUEST"/>
				<item code="4004" name="NOT_FOUND"/>
				<item code="4005" name="OPERATION_NOT_ALLOWED"/>
				<item code="4008" name="REQUEST_TIMEOUT"/>
				<item code="4101" name="SUBSCRIPTION_CREATOR_HAS_NO_PRIVILEGE"/>
				<item code="4102" name="CONTENTS_UNACCEPTABLE"/>
				<item code="4103" name="ORIGINATOR_HAS_NO_PRIVILEGE"/>
				<item code="4104" name="GROUP_REQUEST_IDENTIFIER_EXISTS"/>
				<item code="4105" name="CONFLICT"/>
				<item code="4106" name="ORIGINATOR_HAS_NOT_REGISTERED"/>
				<item code="4107" name="SECURITY_ASSOCIATION_REQUIRED"/>
				<item code="4108" name="INVALID_CHILD_RESOURCE_TYPE"/>
				<item code="4109" name="NO_MEMBERS"/>
				<item code="4110" name="GROUP_MEMBER_TYPE_INCONSISTENT"/>
				<item code="5000" name="INTERNAL_SERVER_ERROR"/>
				<item code="5001" name="NOT_IMPLEMENTED"/>
				<item code="5103" name="TARGET_NOT_REACHABLE"/>
				<item code="5105" name="RECEIVER_HAS_NO_PRIVILEGE"/>
				<item code="5106" name="ALREADY_EXISTS"/>
				<item code="5203" name="TARGET_NOT_SUBSCRIBABLE"/>
				<item code="5204" name="SUBSCRIPTION_VERIFICATION_INITIATION_FAILED"/>
				<item code="5205" name="SUBSCRIPTION_HOST_HAS_NO_PRIVILEGE"/>
				<item code="5206" name="NON_BLOCKING_REQUEST_NOT_SUPPORTED"/>
				<item code="5207" name="NOT_ACCEPTABLE"/>
				<item code="5209" name="GROUP_MEMBERS_NOT_RESPONDED"/>
				<item code="6003" name="EXTERNAL_OBJECT_NOT_REACHABLE"/>
				<item code="6005" name="EXTERNAL_OBJECT_NOT_FOUND"/>
				<item code="6010" name="MAX_NUMBER_OF_MEMBER_EXCEEDED"/>
				<item code="6020" name="MGMT_SESSION_CANNOT_BE_ESTABLISHED"/>
				<item code="6021" name="MGMT_SESSION_ESTABLISHMENT_TIMEOUT"/>
				<item code="6022" name="INVALID_CMDTYPE"/>
				<item code="6023" name="INVALID_ARGUMENTS"/>
				<item code="6024" name="INSUFFICIENT_ARGUMENTS"/>
				<item code="6025" name="MGMT_CONVERSION_ERROR"/>
				<item code="6026" name="MGMT_CANCELLATION_FAILED"/>
				<item code="6028" name="ALREADY_COMPLETE"/>
				<item code="6029" name="MGMT_COMMAND_NOT_CANCELLABLE"/>
			</data>
		</avp>

		<avp name="M2M-Event-Record-Timestamp" code="1010" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Time"/>
		</avp>

		<avp name="Control-Memory-Size" code="1002" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Data-Memory-Size" code="1004" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Access-Network-Identifier" code="1000" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Occupancy" code="1012" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Group-Name" code="1006" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="UTF8String"/>
		</avp>

		<avp name="Maximum-Number-Members" code="1009" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Current-Number-Members" code="1003" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Subgroup-Name" code="1021" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="45687">
			<data type="UTF8String"/>
		</avp>

		<avp name="SCEF-ID" code="3125" must="M,V" may-encrypt="No" vendor-id="10415">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Serving-Node-Identity" code="3929" must="V,M" may="-" vendor-id="10415">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="NIDD-Submission" code="3928" must="V/M" may="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Submission-Timestamp" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Accounting-Input-Octets" required="false" max="1"/>
				<rule avp="Accounting-Output-Octets" required="false" max="1"/>
				<rule avp="Change-Condition" required="false" max="1"/>
			</data>
		</avp>

		<avp name="3GPP-PDP-Type" code="3" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="IPV4"/>
				<item code="1" name="PPP"/>
				<item code="2" name="IPV6"/>
				<item code="3" name="IPV4V6"/>
				<item code="4" name="NON_IP"/>
				<item code="5" name="UNSTRUCTURED"/>
				<item code="6" name="ETHERNET"/>
			</data>
		</avp>
		
		<avp name="Type-Number" code="1204" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="TypeNumber_0"/>
				<item code="1" name="TypeNumber_1"/>
				<item code="2" name="TypeNumber_2"/>
				<item code="3" name="TypeNumber_3"/>
				<item code="4" name="TypeNumber_4"/>
				<item code="5" name="TypeNumber_5"/>
				<item code="6" name="TypeNumber_6"/>
				<item code="7" name="TypeNumber_7"/>
				<item code="8" name="TypeNumber_8"/>
				<item code="9" name="TypeNumber_9"/>
				<item code="10" name="TypeNumber_10"/>
				<item code="12" name="TypeNumber_12"/>
				<item code="13" name="TypeNumber_13"/>
				<item code="14" name="TypeNumber_14"/>
				<item code="15" name="TypeNumber_15"/>
				<item code="16" name="TypeNumber_16"/>
				<item code="17" name="TypeNumber_17"/>
				<item code="18" name="TypeNumber_18"/>
				<item code="19" name="TypeNumber_19"/>
				<item code="20" name="TypeNumber_20"/>
				<item code="21" name="TypeNumber_21"/>
				<item code="22" name="TypeNumber_22"/>
				<item code="23" name="TypeNumber_23"/>
				<item code="24" name="TypeNumber_24"/>
				<item code="25" name="TypeNumber_25"/>
				<item code="26" name="TypeNumber_26"/>
				<item code="27" name="TypeNumber_27"/>
				<item code="28" name="TypeNumber_28"/>
				<item code="29" name="TypeNumber_29"/>
				<item code="30" name="TypeNumber_30"/>
				<item code="31" name="TypeNumber_31"/>
				<item code="32" name="TypeNumber_32"/>
				<item code="33" name="TypeNumber_33"/>
				<item code="34" name="TypeNumber_34"/>
				<item code="35" name="TypeNumber_35"/>
				<item code="36" name="TypeNumber_36"/>
				<item code="37" name="TypeNumber_37"/>
				<item code="38" name="TypeNumber_38"/>
				<item code="39" name="TypeNumber_39"/>
				<item code="40" name="TypeNumber_40"/>
				<item code="41" name="TypeNumber_41"/>
				<item code="42" name="TypeNumber_42"/>
				<item code="43" name="TypeNumber_43"/>
				<item code="44" name="TypeNumber_44"/>
				<item code="45" name="TypeNumber_45"/>
				<item code="46" name="TypeNumber_46"/>
				<item code="47" name="TypeNumber_47"/>
				<item code="48" name="TypeNumber_48"/>
				<item code="49" name="TypeNumber_49"/>
				<item code="50" name="TypeNumber_50"/>
				<item code="51" name="TypeNumber_51"/>
				<item code="52" name="TypeNumber_52"/>
				<item code="53" name="TypeNumber_53"/>
				<item code="54" name="TypeNumber_54"/>
				<item code="55" name="TypeNumber_55"/>
				<item code="56" name="TypeNumber_56"/>
				<item code="58" name="TypeNumber_58"/>
				<item code="59" name="TypeNumber_59"/>
				<item code="60" name="TypeNumber_60"/>
				<item code="61" name="TypeNumber_61"/>
				<item code="62" name="TypeNumber_62"/>
				<item code="63" name="TypeNumber_63"/>
				<item code="64" name="TypeNumber_64"/>
				<item code="65" name="TypeNumber_65"/>
				<item code="66" name="TypeNumber_66"/>
				<item code="67" name="TypeNumber_67"/>
				<item code="68" name="TypeNumber_68"/>
				<item code="69" name="TypeNumber_69"/>
				<item code="70" name="TypeNumber_70"/>
				<item code="71" name="TypeNumber_71"/>
				<item code="72" name="TypeNumber_72"/>
				<item code="73" name="TypeNumber_73"/>
				<item code="74" name="TypeNumber_74"/>
				<item code="75" name="TypeNumber_75"/>
				<item code="76" name="TypeNumber_76"/>
				<item code="77" name="TypeNumber_77"/>
				<item code="78" name="TypeNumber_78"/>
				<item code="79" name="TypeNumber_79"/>
				<item code="80" name="TypeNumber_80"/>
				<item code="81" name="TypeNumber_81"/>
				<item code="82" name="TypeNumber_82"/>
				<item code="83" name="TypeNumber_83"/>
				<item code="84" name="TypeNumber_84"/>
				<item code="85" name="TypeNumber_85"/>
				<item code="86" name="TypeNumber_86"/>
				<item code="87" name="TypeNumber_87"/>
				<item code="88" name="TypeNumber_88"/>
				<item code="90" name="TypeNumber_90"/>
				<item code="91" name="TypeNumber_91"/>
				<item code="513" name="TypeNumber_513"/>
				<item code="514" name="TypeNumber_514"/>
				<item code="515" name="TypeNumber_515"/>
				<item code="516" name="TypeNumber_516"/>
				<item code="517" name="TypeNumber_517"/>
				<item code="518" name="TypeNumber_518"/>
				<item code="519" name="TypeNumber_519"/>
				<item code="520" name="TypeNumber_520"/>
				<item code="521" name="TypeNumber_521"/>
				<item code="522" name="TypeNumber_522"/>
				<item code="523" name="TypeNumber_523"/>
				<item code="524" name="TypeNumber_524"/>
				<item code="768" name="TypeNumber_768"/>
				<item code="769" name="TypeNumber_769"/>
				<item code="770" name="TypeNumber_770"/>
				<item code="773" name="TypeNumber_773"/>
				<item code="774" name="TypeNumber_774"/>
				<item code="775" name="TypeNumber_775"/>
				<item code="776" name="TypeNumber_776"/>
				<item code="777" name="TypeNumber_777"/>
				<item code="778" name="TypeNumber_778"/>
				<item code="779" name="TypeNumber_779"/>
				<item code="780" name="TypeNumber_780"/>
				<item code="781" name="TypeNumber_781"/>
				<item code="782" name="TypeNumber_782"/>
				<item code="783" name="TypeNumber_783"/>
				<item code="784" name="TypeNumber_784"/>
				<item code="785" name="TypeNumber_785"/>
				<item code="786" name="TypeNumber_786"/>
				<item code="787" name="TypeNumber_787"/>
				<item code="788" name="TypeNumber_788"/>
				<item code="789" name="TypeNumber_789"/>
				<item code="790" name="TypeNumber_790"/>
				<item code="791" name="TypeNumber_791"/>
				<item code="792" name="TypeNumber_792"/>
				<item code="793" name="TypeNumber_793"/>
				<item code="794" name="TypeNumber_794"/>
				<item code="795" name="TypeNumber_795"/>
				<item code="796" name="TypeNumber_796"/>
				<item code="797" name="TypeNumber_797"/>
				<item code="798" name="TypeNumber_798"/>
			</data>
		</avp>
	
		<avp name="Application-Service-Type" code="2102" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
			<data type="Enumerated">
				<item code="100" name="SIMPLE_IM_SENDING"/>
				<item code="101" name="RECEIVING"/>
				<item code="102" name="RETRIEVAL"/>
				<item code="103" name="INVITING"/>
				<item code="104" name="LEAVING"/>
				<item code="105" name="JOINING"/>
			</data>
		</avp>

		<avp name="AoC-Request-Type" code="2055" must="V,M" may="-" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="AOC_NOT_REQUESTED"/>
				<item code="1" name="AOC_FULL"/>
				<item code="2" name="AOC_COST_ONLY"/>
				<item code="3" name="AOC_TARIFF_ONLY"/>
			</data>
		</avp>

		<avp name="APN-Aggregate-Max-Bitrate-DL" code="1040" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="APN-Aggregate-Max-Bitrate-UL" code="1041" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Bearer-Identifier" code="1020" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="CC-Request-Number" code="415" must="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.2 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="CC-Request-Type" code="416" must="M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.3 -->
			<data type="Enumerated">
				<item code="1" name="INITIAL_REQUEST"/>
				<item code="2" name="UPDATE_REQUEST"/>
				<item code="3" name="TERMINATION_REQUEST"/>
				<item code="4" name="EVENT_REQUEST"/>
			</data>
		</avp>

		<avp name="Conditional-APN-Aggregate-Max-Bitrate" code="2818" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Grouped">
				<rule avp="APN-Aggregate-Max-Bitrate-UL" required="false" max="1"/>
				<rule avp="APN-Aggregate-Max-Bitrate-DL" required="false" max="1"/>
				<rule avp="Extended-APN-AMBR-UL" required="false" max="1"/>
				<rule avp="Extended-APN-AMBR-DL" required="false" max="1"/>
				<rule avp="IP-CAN-Type" required="false"/>
				<rule avp="RAT-Type" required="false"/>
			</data>
		</avp>

		<avp name="Content-Version" code="552" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned64"/>
		</avp>

		<avp name="Event-Charging-TimeStamp" code="1258" must="V,M" may="-" vendor-id="10415">
			<data type="Time"/>
		</avp>

		<avp name="Extended-APN-AMBR-DL" code="2848" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Extended-APN-AMBR-UL" code="2849" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Extended-GBR-DL" code="2850" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Extended-GBR-UL" code="2851" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Feature-List" code="630" must="V" must-not="M" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Feature-List-ID" code="629" must="V" must-not="M" may-encrypt="No" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Filter-Id" code="11" must="M" may="-" must-not="V" vendor-id="0">
			<data type="UTF8String"/>
		</avp>

		<avp name="Guaranteed-Bitrate-DL" code="1025" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Guaranteed-Bitrate-UL" code="1026" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Media-Component-Status" code="549" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Origin-State-Id" code="278" must="M" may="-" must-not="V" vendor-id="0">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Presence-Reporting-Area-Elements-List" code="2820" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Presence-Reporting-Area-Identifier" code="2821" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="OctetString"/>
		</avp>

		<avp name="Presence-Reporting-Area-Node" code="2855" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
			<data type="Unsigned32"/>
		</avp>

		<avp name="PS-Append-Free-Format-Data" code="867" must="V,M" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="Append"/>
				<item code="0" name="Overwrite"/>
			</data>
		</avp>

		<avp name="Redirect-Server" code="434" must="M" may="-" must-not="V" vendor-id="0">
			<data type="Grouped">
				<rule avp="Redirect-Address-Type" required="true" max="1"/>
				<rule avp="Redirect-Server-Address" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Service-Information" code="873" must="V,M" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Subscription-Id" required="false"/>
				<rule avp="AoC-Information" required="false" max="1"/>
				<rule avp="PS-Information" required="false" max="1"/>
				<rule avp="IMS-Information" required="false" max="1"/>
				<rule avp="MMS-Information" required="false" max="1"/>
				<rule avp="LCS-Information" required="false" max="1"/>
				<rule avp="PoC-Information" required="false" max="1"/>
				<rule avp="MBMS-Information" required="false" max="1"/>
				<rule avp="SMS-Information" required="false" max="1"/>
				<rule avp="VCS-Information" required="false" max="1"/>
				<rule avp="MMTel-Information" required="false" max="1"/>
				<rule avp="ProSe-Information" required="false" max="1"/>
				<rule avp="Service-Generic-Information" required="false" max="1"/>
				<rule avp="IM-Information" required="false" max="1"/>
				<rule avp="DCD-Information" required="false" max="1"/>
				<rule avp="M2M-Information" required="false" max="1"/>
				<rule avp="CPDT-Information" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Low-Balance-Indication" code="2020" must="V,M" may="P" must-not="-" may-encrypt="N" vendor-id="10415">
			<data type="Enumerated">
				<item code="0" name="NOT-APPLICABLE"/>
				<item code="1" name="YES"/>
			</data>
		</avp>

		<avp name="Remaining-Balance" code="2021" must="V,M" may="-" must-not="-" vendor-id="10415">
			<data type="Grouped">
				<rule avp="Unit-Value" required="true" max="1"/>
				<rule avp="Currency-Code" required="true" max="1"/>
			</data>
		</avp>

		<avp name="NAS-Identifier" code="32" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4005#section-9.3.1 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="NAS-IP-Address" code="4" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4005#section-9.3.2 -->
			<data type="OctetString"/>
		</avp>

		<avp name="NAS-IPv6-Address" code="95" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4005#section-9.3.3 -->
			<data type="OctetString"/>
		</avp>

		<avp name="Origin-AAA-Protocol" code="408" must="M" may="-" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4005#section-9.3.6 -->
			<data type="Enumerated">
				<item code="1" name="RADIUS"/>
			</data>
		</avp>

		<avp name="Load-Type" code="651" must="-" may="-" must-not="-" may-encrypt="-">
			<!-- https://datatracker.ietf.org/doc/html/rfc8583#section-7.2 -->
			<data type="Enumerated">
				<item code="0" name="HOST"/>
				<item code="1" name="PEER"/>
			</data>
		</avp>

		<avp name="Load-Value" code="652" must="-" must-not="V" may-encrypt="No">
			<!-- https://datatracker.ietf.org/doc/html/rfc8583#section-7.3 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="SourceID" code="649" must="-" must-not="V" may-encrypt="No">
			<!-- https://datatracker.ietf.org/doc/html/rfc8583#section-7.4 -->
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Packet-Filter-Content" code="1059" must="M,V" must-not="-" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.54 -->
            <data type="OctetString"/>
        </avp>

	</application>
</diameter>`
