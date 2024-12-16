package base

var DefConnBaseDictionary = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

	<application id="0" name="Base"> <!-- Diameter peer connection Common Messages -->
		<!-- Diameter Credit Control Application -->

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

		<avp name="Host-IP-Address" code="257" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Address"/>
		</avp>

		<avp name="Auth-Application-Id" code="258" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Acct-Application-Id" code="259" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Vendor-Specific-Application-Id" code="260" must="-" may="-" must-not="-" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Vendor-Id" required="false" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Origin-Host" code="264" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Supported-Vendor-Id" code="265" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Vendor-Id" code="266" must="-" may="-" must-not="-" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Firmware-Revision" code="267" must="-" may="-" must-not="P,V,M" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Result-Code" code="268" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Product-Name" code="269" must="-" may="-" must-not="P,V,M" may-encrypt="-">
			<data type="UTF8String"/>
		</avp>

		<avp name="Disconnect-Cause" code="273" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="REBOOTING"/>
				<item code="1" name="BUSY"/>
				<item code="2" name="DO_NOT_WANT_TO_TALK_TO_YOU"/>
			</data>
		</avp>

		<avp name="Origin-State-Id" code="278" must="M" may="-" must-not="V" vendor-id="0">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Failed-AVP" code="279" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="Grouped"/>
		</avp>

		<avp name="Error-Message" code="281" must="-" may="-" must-not="-" may-encrypt="-">
			<data type="UTF8String"/>
		</avp>

		<avp name="Origin-Realm" code="296" must="M" may="-" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

    </application>
</diameter>`
