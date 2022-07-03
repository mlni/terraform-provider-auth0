package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGuardian(t *testing.T) {
	httpRecorder := configureHTTPRecorder(t)

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviders(httpRecorder),
		Steps: []resource.TestStep{
			{
				Config: testAccConfigureTwilio,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.message_types.0", "sms"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.provider", "twilio"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.enrollment_message", "enroll foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.verification_message", "verify foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.from", "from bar"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.messaging_service_sid", "foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.auth_token", "bar"),
				),
			},
			{
				Config: testAccConfigureTwilioUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.#", "0"),
				),
			},
			{
				Config: testAccConfigureTwilio,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.message_types.0", "sms"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.provider", "twilio"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.enrollment_message", "enroll foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.verification_message", "verify foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.from", "from bar"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.messaging_service_sid", "foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.auth_token", "bar"),
				),
			},
			{
				Config: testAccConfigureCustomPhone,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.message_types.0", "sms"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.provider", "phone-message-hook"),
				),
			},
			{
				Config: testAccConfigureAuth0Custom,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.message_types.0", "voice"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.provider", "auth0"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.enrollment_message", "enroll foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.verification_message", "verify foo"),
				),
			},
			{
				Config: testAccConfigureAuth0,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.message_types.0", "voice"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.provider", "auth0"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.enrollment_message", "enroll foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.verification_message", "verify foo"),
				),
			},
			{
				Config: testAccConfigureNoPhone,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.#", "0"),
				),
			},
			{
				Config: testAccConfigureEmail,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "email", "true"),
				),
			},
			{
				Config: testAccConfigureEmailUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "email", "false"),
				),
			},
			{
				Config: testAccConfigureOTP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "otp", "true"),
				),
			},
			{
				Config: testAccConfigureOTPUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "otp", "false"),
				),
			},
		},
	})
}

const testAccConfigureAuth0Custom = `
resource "auth0_guardian" "foo" {
  policy = "all-applications"
  phone {
    provider = "auth0"
    message_types = ["voice"]
    options {
      enrollment_message = "enroll foo"
      verification_message = "verify foo"
    }
  }
}
`

const testAccConfigureCustomPhone = `
resource "auth0_guardian" "foo" {
  policy = "all-applications"
  phone {
	provider = "phone-message-hook"
	message_types = ["sms"]
	options{
	}
  }
}
`
const testAccConfigureTwilio = `
resource "auth0_guardian" "foo" {
  policy = "all-applications"
  phone {
	provider = "twilio"
	message_types = ["sms"]
	options {
		enrollment_message = "enroll foo"
		verification_message = "verify foo"
		from = "from bar"
		messaging_service_sid = "foo"
		auth_token = "bar"
		sid = "foo"
	}
  }
}
`

const testAccConfigureTwilioUpdate = `
resource "auth0_guardian" "foo" {
  policy = "all-applications"
}
`

const testAccConfigureAuth0 = `
resource "auth0_guardian" "foo" {
  policy = "all-applications"
  phone {
	provider = "auth0"
	message_types = ["voice"]
	options {
		enrollment_message = "enroll foo"
		verification_message = "verify foo"
	}
}
}
`

const testAccConfigureNoPhone = `
resource "auth0_guardian" "foo" {
  policy = "all-applications"
}
`

const testAccConfigureEmail = `
resource "auth0_guardian" "foo" {
  policy = "all-applications"
  email  = true
}
`

const testAccConfigureEmailUpdate = `
resource "auth0_guardian" "foo" {
  policy = "all-applications"
  email  = false
}
`

const testAccConfigureOTP = `
resource "auth0_guardian" "foo" {
	policy = "all-applications"
	otp  = true
}
`

const testAccConfigureOTPUpdate = `
resource "auth0_guardian" "foo" {
	policy = "all-applications"
	otp  = false
}
`

func TestAccGuardianPhoneMessageHookWithNoOptions(t *testing.T) {
	httpRecorder := configureHTTPRecorder(t)

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviders(httpRecorder),
		Steps: []resource.TestStep{
			{
				Config: testAccGuardianPhoneMessageHookWithNoOptions,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.default", "phone.#", "0"),
				),
			},
			{
				Config: testAccGuardianPhoneMessageHookWithNoOptionsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.default", "phone.#", "1"),
					resource.TestCheckResourceAttr("auth0_guardian.default", "phone.0.provider", "phone-message-hook"),
					resource.TestCheckResourceAttr("auth0_guardian.default", "phone.0.message_types.0", "sms"),
					resource.TestCheckResourceAttr("auth0_guardian.default", "phone.0.options.#", "1"),
				),
			},
			{
				Config: testAccGuardianPhoneMessageHookWithNoOptionsUpdate2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.default", "phone.#", "1"),
					resource.TestCheckResourceAttr("auth0_guardian.default", "phone.0.provider", "phone-message-hook"),
					resource.TestCheckResourceAttr("auth0_guardian.default", "phone.0.message_types.0", "sms"),
					resource.TestCheckResourceAttr("auth0_guardian.default", "phone.0.options.#", "1"),
				),
			},
		},
	})
}

const testAccGuardianPhoneMessageHookWithNoOptions = `
resource "auth0_guardian" "default" {
  policy     = "all-applications"
  otp        = false
  email      = false
}
`

const testAccGuardianPhoneMessageHookWithNoOptionsUpdate = `
resource "auth0_guardian" "default" {
  policy     = "all-applications"
  otp        = false
  email      = false
  phone {
    provider = "phone-message-hook"
    message_types = ["sms"]
  }
}
`

const testAccGuardianPhoneMessageHookWithNoOptionsUpdate2 = `
resource "auth0_guardian" "default" {
  policy     = "all-applications"
  otp        = false
  email      = false
  phone {
    provider = "phone-message-hook"
    message_types = ["sms"]
	options {}
  }
}
`

func TestAccGuardianWebAuthnRoaming(t *testing.T) {
	httpRecorder := configureHTTPRecorder(t)

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviders(httpRecorder),
		Steps: []resource.TestStep{
			{
				Config: testAccConfigureWebAuthnRoamingCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "webauthn_roaming.#", "1"),
				),
			},
			{
				Config: testAccConfigureWebAuthnRoamingUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "webauthn_roaming.#", "1"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "webauthn_roaming.0.user_verification", "required"),
				),
			},
			{
				Config: testAccConfigureWebAuthnRoamingDelete,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "webauthn_roaming.#", "0"),
				),
			},
		},
	})
}

const testAccConfigureWebAuthnRoamingCreate = `
resource "auth0_guardian" "foo" {
	policy = "all-applications"
	webauthn_roaming {}
}
`

const testAccConfigureWebAuthnRoamingUpdate = `
resource "auth0_guardian" "foo" {
	policy = "all-applications"
	webauthn_roaming {
		user_verification = "required"
	}
}
`

const testAccConfigureWebAuthnRoamingDelete = `
resource "auth0_guardian" "foo" {
	policy = "all-applications"
}
`

func TestAccGuardianWebAuthnPlatform(t *testing.T) {
	httpRecorder := configureHTTPRecorder(t)

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviders(httpRecorder),
		Steps: []resource.TestStep{
			{
				Config: testAccConfigureWebAuthnPlatformCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "webauthn_platform.#", "1"),
				),
			},
			{
				Config: testAccConfigureWebAuthnPlatformDelete,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "webauthn_platform.#", "0"),
				),
			},
		},
	})
}

const testAccConfigureWebAuthnPlatformCreate = `
resource "auth0_guardian" "foo" {
	policy = "all-applications"
	webauthn_platform {}
}
`

const testAccConfigureWebAuthnPlatformDelete = `
resource "auth0_guardian" "foo" {
	policy = "all-applications"
}
`

func TestAccGuardianDUO(t *testing.T) {
	httpRecorder := configureHTTPRecorder(t)

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviders(httpRecorder),
		Steps: []resource.TestStep{
			{
				Config: testAccConfigureDUOCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "duo.#", "1"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "duo.0.hostname", "api-hostname"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "duo.0.secret_key", "someSecret"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "duo.0.integration_key", "someKey"),
				),
			},
			{
				Config: testAccConfigureDUODelete,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "duo.#", "0"),
				),
			},
		},
	})
}

const testAccConfigureDUOCreate = `
resource "auth0_guardian" "foo" {
	policy = "all-applications"
	duo {
		integration_key = "someKey"
		secret_key = "someSecret"
		hostname = "api-hostname"
	}
}
`

const testAccConfigureDUODelete = `
resource "auth0_guardian" "foo" {
	policy = "all-applications"
}
`
