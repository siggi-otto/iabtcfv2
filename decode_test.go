package iabtcfv2

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	version, err := GetVersion("BOr70tQOxPQw-BcAsCFRDEqAAAAu1rxyZn7kfUXiXSZxNuiGGp6h-Wd9CWUcKZYpMAnyhYZRfg_AQhQ4Eu0LRNNycgh45MoCCMoRQaiSkCABGgFcTpjTmxAUxoRLawAMBrwhWLEQeroyHcJzAAHN_QjACAA")
	if err != nil {
		t.Errorf("Version should be decoded without error: %s", err)
		return
	}

	if version != TcfVersion1 {
		t.Errorf("Version should be %d", TcfVersion1)
	}

	version, err = GetVersion("COy7f9HOy7f_1BcABBENAjCoAPKAAFKAAAqIDaQCQABAAVAAyACAAFoANQAkgEdANoA2kAYAAQAFQAMgAgABaAbQAUMAQAAEABUADIAIAAWgBJgDCAMQA9ACEAEdAKuAXUAwIBhADRAG0FAEQABAAVAAyACAAFoANQAkwBhAGIAegBCACOgFXALqAYEAwgBogDaDACIAAgAKgAZABAAC0AGoASYAwgDEAPQAhABHQCrgF1AMCAYQA0QBtCABEAAQAFQAMgAgABaADUAJMAYQBiAHoAQgAjoBVwC6gGBAMIAaIA2hQAiAAIACoAGQAQAAtABqAEmAMIAxAD0AIQAR0Aq4BdQDAgGEANEAbQA.cAAACAAAAUg")
	if err != nil {
		t.Errorf("Version should be decoded without error: %s", err)
		return
	}

	if version != TcfVersion2 {
		t.Errorf("Version should be %d", TcfVersion2)
	}
}

func TestGetSegmentType(t *testing.T) {
	str := "IF3EXySoGY2tho2YVFzBEIYwfJxyigMgShgQIsS0NQIeFLBoGPiAAHBGYJAQAGBAkkACBAQIsHGBMCQABgAgRiRCMQEGMDzNIBIBAggkbY0FACCVmnkHS3ZCY70"

	segType, err := GetSegmentType(str)
	if err != nil {
		t.Errorf("Segment type should be decoded without error: %s", err)
		return
	}

	if segType != SegmentTypeDisclosedVendors {
		t.Errorf("Segment type should be %d", SegmentTypeDisclosedVendors)
	}
}

func TestGetVersionAndSegmentTypeFail(t *testing.T) {
	str := "A"

	_, err := GetVersion(str)
	if err == nil {
		t.Errorf("Version should not be decoded")
		return
	}

	_, err = GetSegmentType(str)
	if err == nil {
		t.Errorf("Segment type should not be decoded")
	}
}

func TestDecode(t *testing.T) {
	str := "COxR03kOxR1CqBcABCENAgCMAP_AAH_AAAqIF3EXySoGY2thI2YVFxBEIYwfJxyigMgChgQIsSwNQIeFLBoGLiAAHBGYJAQAGBAEEACBAQIkHGBMCQAAgAgBiRCMQEGMCzNIBIBAggEbY0FACCVmHkHSmZCY7064O__QLuIJEFQMAkSBAIACLECIQwAQDiAAAYAlAAABAhIaAAgIWBQEeAAAACAwAAgAAABBAAACAAQAAICIAAABAAAgAiAQAAAAGgIQAACBABACRIAAAEANCAAgiCEAQg4EAo4AAA.IF3EXySoGY2tho2YVFzBEIYwfJxyigMgShgQIsS0NQIeFLBoGPiAAHBGYJAQAGBAkkACBAQIsHGBMCQABgAgRiRCMQEGMDzNIBIBAggkbY0FACCVmnkHS3ZCY70-6u__QA.elAAAAAAAWA"

	data, err := Decode(str)
	if err != nil {
		t.Errorf("TC String should be decoded without error: %s", err)
		return
	}

	result := data.ToTCString()
	if result == "" {
		t.Errorf("Encode() should be produce a string")
		return
	}

	if result != str {
		t.Errorf("Encode() should produce the same string: in = %s, out = %s", str, result)
	}
}

func TestDecodeMissingCore(t *testing.T) {
	str := "IF3EXySoGY2tho2YVFzBEIYwfJxyigMgShgQIsS0NQIeFLBoGPiAAHBGYJAQAGBAkkACBAQIsHGBMCQABgAgRiRCMQEGMDzNIBIBAggkbY0FACCVmnkHS3ZCY70-6u__QA.elAAAAAAAWA"

	_, err := Decode(str)
	if err == nil {
		t.Errorf("TC String should not be decoded: %s", err)
		return
	}
}

func TestDecodeWrongOrdered(t *testing.T) {
	str := "elAAAAAAAWA.COxR03kOxR1CqBcABCENAgCMAP_AAH_AAAqIF3EXySoGY2thI2YVFxBEIYwfJxyigMgChgQIsSwNQIeFLBoGLiAAHBGYJAQAGBAEEACBAQIkHGBMCQAAgAgBiRCMQEGMCzNIBIBAggEbY0FACCVmHkHSmZCY7064O__QLuIJEFQMAkSBAIACLECIQwAQDiAAAYAlAAABAhIaAAgIWBQEeAAAACAwAAgAAABBAAACAAQAAICIAAABAAAgAiAQAAAAGgIQAACBABACRIAAAEANCAAgiCEAQg4EAo4AAA.IF3EXySoGY2tho2YVFzBEIYwfJxyigMgShgQIsS0NQIeFLBoGPiAAHBGYJAQAGBAkkACBAQIsHGBMCQABgAgRiRCMQEGMDzNIBIBAggkbY0FACCVmnkHS3ZCY70-6u__QA"

	_, err := Decode(str)
	if err == nil {
		t.Errorf("TC String should not be decoded: %s", err)
		return
	}
}

func TestDecodeDuplicateSegment(t *testing.T) {
	str := "COxR03kOxR1CqBcABCENAgCMAP_AAH_AAAqIF3EXySoGY2thI2YVFxBEIYwfJxyigMgChgQIsSwNQIeFLBoGLiAAHBGYJAQAGBAEEACBAQIkHGBMCQAAgAgBiRCMQEGMCzNIBIBAggEbY0FACCVmHkHSmZCY7064O__QLuIJEFQMAkSBAIACLECIQwAQDiAAAYAlAAABAhIaAAgIWBQEeAAAACAwAAgAAABBAAACAAQAAICIAAABAAAgAiAQAAAAGgIQAACBABACRIAAAEANCAAgiCEAQg4EAo4AAA.COxR03kOxR1CqBcABCENAgCMAP_AAH_AAAqIF3EXySoGY2thI2YVFxBEIYwfJxyigMgChgQIsSwNQIeFLBoGLiAAHBGYJAQAGBAEEACBAQIkHGBMCQAAgAgBiRCMQEGMCzNIBIBAggEbY0FACCVmHkHSmZCY7064O__QLuIJEFQMAkSBAIACLECIQwAQDiAAAYAlAAABAhIaAAgIWBQEeAAAACAwAAgAAABBAAACAAQAAICIAAABAAAgAiAQAAAAGgIQAACBABACRIAAAEANCAAgiCEAQg4EAo4AAA"

	_, err := Decode(str)
	if err == nil {
		t.Errorf("TC String should not be decoded: %s", err)
		return
	}
}

func TestDecodeCoreString(t *testing.T) {
	str := "COxR03kOxR1CqBcABCENAgCMAP_AAH_AAAqIF3EXySoGY2thI2YVFxBEIYwfJxyigMgChgQIsSwNQIeFLBoGLiAAHBGYJAQAGBAEEACBAQIkHGBMCQAAgAgBiRCMQEGMCzNIBIBAggEbY0FACCVmHkHSmZCY7064O__QLuIJEFQMAkSBAIACLECIQwAQDiAAAYAlAAABAhIaAAgIWBQEeAAAACAwAAgAAABBAAACAAQAAICIAAABAAAgAiAQAAAAGgIQAACBABACRIAAAEANCAAgiCEAQg4EAo4AAA"

	segType, err := GetSegmentType(str)
	if err != nil {
		t.Errorf("Segment type should be decoded without error: %s", err)
		return
	}

	if segType != SegmentTypeCoreString {
		t.Errorf("Segment type should be %d", SegmentTypeCoreString)
		return
	}

	segment, err := DecodeCoreString(str)
	if err != nil {
		t.Errorf("Segment should be decoded without error: %s", err)
		return
	}

	result := segment.Encode()
	if result == "" {
		t.Errorf("Encode() should be produce a string")
		return
	}

	if result != str {
		t.Errorf("Encode() should produce the same string: in = %s, out = %s", str, result)
	}
}

func TestDecodeDisclosedVendors(t *testing.T) {
	str := "IF3EXySoGY2tho2YVFzBEIYwfJxyigMgShgQIsS0NQIeFLBoGPiAAHBGYJAQAGBAkkACBAQIsHGBMCQABgAgRiRCMQEGMDzNIBIBAggkbY0FACCVmnkHS3ZCY70-6u__QA"

	segType, err := GetSegmentType(str)
	if err != nil {
		t.Errorf("Segment type should be decoded without error: %s", err)
		return
	}

	if segType != SegmentTypeDisclosedVendors {
		t.Errorf("Segment type should be %d", SegmentTypeDisclosedVendors)
		return
	}

	segment, err := DecodeDisclosedVendors(str)
	if err != nil {
		t.Errorf("Segment should be decoded without error: %s", err)
		return
	}

	if segment.IsVendorDisclosed(1) {
		t.Errorf("Vendor 1 should not be disclosed")
		return
	}

	if !segment.IsVendorDisclosed(9) {
		t.Errorf("Vendor 9 should be disclosed")
		return
	}

	result := segment.Encode()
	if result == "" {
		t.Errorf("Encode() should be produce a string")
		return
	}

	if result != str {
		t.Errorf("Encode() should produce the same string: in = %s, out = %s", str, result)
	}
}

func TestDecodeAllowedVendors(t *testing.T) {
	str := "QF3QAgABAA1A"

	segType, err := GetSegmentType(str)
	if err != nil {
		t.Errorf("Segment type should be decoded without error: %s", err)
		return
	}

	if segType != SegmentTypeAllowedVendors {
		t.Errorf("Segment type should be %d", SegmentTypeAllowedVendors)
		return
	}

	segment, err := DecodeAllowedVendors(str)
	if err != nil {
		t.Errorf("Segment should be decoded without error: %s", err)
		return
	}

	if segment.IsVendorAllowed(10) {
		t.Errorf("Vendor 10 should not be disclosed")
		return
	}

	if !segment.IsVendorAllowed(53) {
		t.Errorf("Vendor 53 should be disclosed")
		return
	}

	result := segment.Encode()
	if result == "" {
		t.Errorf("Encode() should be produce a string")
		return
	}

	if result != str {
		t.Errorf("Encode() should produce the same string: in = %s, out = %s", str, result)
	}
}

func TestDecodePublisherTC(t *testing.T) {
	str := "elAAAAAAAWA"

	segType, err := GetSegmentType(str)
	if err != nil {
		t.Errorf("Segment type should be decoded without error: %s", err)
		return
	}

	if segType != SegmentTypePublisherTC {
		t.Errorf("Segment type should be %d", SegmentTypePublisherTC)
		return
	}

	segment, err := DecodePublisherTC(str)
	if err != nil {
		t.Errorf("Segment should be decoded without error: %s", err)
		return
	}

	if !segment.IsPurposeAllowed(1) {
		t.Errorf("Purpose 1 should be allowed")
		return
	}

	if segment.NumCustomPurposes != 2 {
		t.Errorf("NumCustomPurposes should be 2")
	}

	result := segment.Encode()
	if result == "" {
		t.Errorf("Encode() should be produce a string")
		return
	}

	if result != str {
		t.Errorf("Encode() should produce the same string: in = %s, out = %s", str, result)
	}
}

func TestPublisherRestrictionFlexibleVendor(t *testing.T) {
	str := "CPStgrQPStgrQAGABCDEB9CsAP_AAH_AAAqIH-NN7S__a2Pj-359Q_t0eY1f9953v-UhjhaZk6QF0bPDsL8V4mM6vE3opioKuBYEO3LAIQRlHKHcBQGAaokRoTPsbk2MLpAAJ7PEmgMbEmdIGHV9m93DnZKYz3w-2r6T_u4NRP_M5MfpP41v3Wt5tl06qXTTVz8YhLP1cAABAAAAQPiAIEBAUAgAEMAEQAFCIQAAQpiQAAAABBCABAAAAIiAAQVwAZIIEAAARAAAQAABAQgwAAAAAABCAAAACwQCAACAQAAgAEAAAAEJAIBACAEAAAEAJABACACECAggAAAwDAgAACCABABAAACJDAAAMIIASABgBEAABEgAGAAACAoMgFgBMAEcAMsAfYBWwExAJsAWwAz4BygD4hEAkAZYBTwDqgHyAQ6AkQBNgDPgHKCQAIDfxAAEAEgSBUAAgABYAFQAMgAcAA8ACAAGUANAA1AB5AEQARQAmABvADmAHoAP0AiACJAEsAJoAUoAtwBhwDKAMsAaoA-wB-gEUAKeAbQA3AB8gEOgJEATEAmwBTQC2AGSAM-AaQA1iByYHKBQAYAigBfAO3CAAwASAGiAU-GgGgBcAGWAQUAp8BaAFpAOqAfIBDoCRAE2AMYAZ8A5QOABAb-KgGABMAC4AI4AZcBaAFpASCAmIBNgCmwFsAM-AcoOgZAALAAqABkADgAIIAYgBlADQANQAeAA-gCIAIoATAAuABiADMAG8AOYAegA_ACIAEsAJgATQAowBSgC3AGGAMoAaIA-wB-gEUAKfAWgBaQC8gG4AOoAh0BIICRAE2AKagWwBbIDGAGSAMsAZmAz4BpADWIHJgcoPADAAqAEUAL4AjIDfwHbjgAIAJCEBYABYAGQAYgBMAC4AGIAMwAbwA9ACOAH2ARQAoYBT4C0ALSAdQBIICRAE2AKagWwBbIDPiIAMAFQAvgCMkoEAACAAFgAZAA4AB8AGIAPAAiABMAC4AGIAMwAbYBEAESAKMAUoAtwBqgEnAKfAWgBaQDcAHUAPkAh0BIgCbAFsAMsAZ8A0gBrBMAEARkBv5SBQAAsACoAGQAOAAggBiAGUANAA1AB5AEQARQAmABSADEAGYAOYAfgBEACjAFKALcAZQA0QBqgD7AKGAVsAvIBtADcAIdASIAk4BNgC2AGMAMkAZYAz4BpADWIHJgcoVACAAqAB8AL4Bv5QAGACQAk4BOw.YAAAAAAAAAAA"
	tcData, err := Decode(str)
	if err != nil {
		t.Errorf("fail to decode tcstring: %s", err)
		return
	}
	if !tcData.IsVendorAllowedForFlexiblePurposesLI(916, 4) {
		t.Errorf("flexible vendor 916 should be allowed to purpose 4 even with publisher restriction because consent is established on purpose 4 for this vendor")
	}
}
