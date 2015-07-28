package otr3

import "testing"

func Test_policies_allowV1_addsAllowingOfV1(t *testing.T) {
	p := policies(0)
	p.allowV1()
	assertEquals(t, p.has(allowV1), true)
}

func Test_policies_requireEncryption_addsRequirementOfEncryption(t *testing.T) {
	p := policies(0)
	p.requireEncryption()
	assertEquals(t, p.has(requireEncryption), true)
}

func Test_policies_sendWhitespaceTag_addsPolicyForSendingWhitespaceTag(t *testing.T) {
	p := policies(0)
	p.sendWhitespaceTag()
	assertEquals(t, p.has(sendWhitespaceTag), true)
}

func Test_policies_whitespaceStartAKE_addsWhitespaceStartAKEPolicy(t *testing.T) {
	p := policies(0)
	p.whitespaceStartAKE()
	assertEquals(t, p.has(whitespaceStartAke), true)
}

func Test_policies_errorStartAKE_addsErrorStartAKEPolicy(t *testing.T) {
	p := policies(0)
	p.errorStartAKE()
	assertEquals(t, p.has(errorStartAke), true)
}