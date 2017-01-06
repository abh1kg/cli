// This file was generated by counterfeiter
package commandfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/util/configv3"
)

type FakeConfig struct {
	APIVersionStub        func() string
	aPIVersionMutex       sync.RWMutex
	aPIVersionArgsForCall []struct{}
	aPIVersionReturns     struct {
		result1 string
	}
	AccessTokenStub        func() string
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct{}
	accessTokenReturns     struct {
		result1 string
	}
	BinaryBuildDateStub        func() string
	binaryBuildDateMutex       sync.RWMutex
	binaryBuildDateArgsForCall []struct{}
	binaryBuildDateReturns     struct {
		result1 string
	}
	BinaryNameStub        func() string
	binaryNameMutex       sync.RWMutex
	binaryNameArgsForCall []struct{}
	binaryNameReturns     struct {
		result1 string
	}
	BinaryVersionStub        func() string
	binaryVersionMutex       sync.RWMutex
	binaryVersionArgsForCall []struct{}
	binaryVersionReturns     struct {
		result1 string
	}
	ColorEnabledStub        func() configv3.ColorSetting
	colorEnabledMutex       sync.RWMutex
	colorEnabledArgsForCall []struct{}
	colorEnabledReturns     struct {
		result1 configv3.ColorSetting
	}
	CurrentUserStub        func() (configv3.User, error)
	currentUserMutex       sync.RWMutex
	currentUserArgsForCall []struct{}
	currentUserReturns     struct {
		result1 configv3.User
		result2 error
	}
	DialTimeoutStub        func() time.Duration
	dialTimeoutMutex       sync.RWMutex
	dialTimeoutArgsForCall []struct{}
	dialTimeoutReturns     struct {
		result1 time.Duration
	}
	ExperimentalStub        func() bool
	experimentalMutex       sync.RWMutex
	experimentalArgsForCall []struct{}
	experimentalReturns     struct {
		result1 bool
	}
	LocaleStub        func() string
	localeMutex       sync.RWMutex
	localeArgsForCall []struct{}
	localeReturns     struct {
		result1 string
	}
	MinCLIVersionStub        func() string
	minCLIVersionMutex       sync.RWMutex
	minCLIVersionArgsForCall []struct{}
	minCLIVersionReturns     struct {
		result1 string
	}
	PluginsStub        func() map[string]configv3.Plugin
	pluginsMutex       sync.RWMutex
	pluginsArgsForCall []struct{}
	pluginsReturns     struct {
		result1 map[string]configv3.Plugin
	}
	RefreshTokenStub        func() string
	refreshTokenMutex       sync.RWMutex
	refreshTokenArgsForCall []struct{}
	refreshTokenReturns     struct {
		result1 string
	}
	PollingIntervalStub        func() time.Duration
	pollingIntervalMutex       sync.RWMutex
	pollingIntervalArgsForCall []struct{}
	pollingIntervalReturns     struct {
		result1 time.Duration
	}
	OverallPollingTimeoutStub        func() time.Duration
	overallPollingTimeoutMutex       sync.RWMutex
	overallPollingTimeoutArgsForCall []struct{}
	overallPollingTimeoutReturns     struct {
		result1 time.Duration
	}
	SetAccessTokenStub        func(token string)
	setAccessTokenMutex       sync.RWMutex
	setAccessTokenArgsForCall []struct {
		token string
	}
	SetOrganizationInformationStub        func(guid string, name string)
	setOrganizationInformationMutex       sync.RWMutex
	setOrganizationInformationArgsForCall []struct {
		guid string
		name string
	}
	SetSpaceInformationStub        func(guid string, name string, allowSSH bool)
	setSpaceInformationMutex       sync.RWMutex
	setSpaceInformationArgsForCall []struct {
		guid     string
		name     string
		allowSSH bool
	}
	UnsetSpaceInformationStub        func()
	unsetSpaceInformationMutex       sync.RWMutex
	unsetSpaceInformationArgsForCall []struct{}
	SetRefreshTokenStub              func(token string)
	setRefreshTokenMutex             sync.RWMutex
	setRefreshTokenArgsForCall       []struct {
		token string
	}
	SetTargetInformationStub        func(api string, apiVersion string, auth string, loggregator string, minCLIVersion string, doppler string, uaa string, routing string, skipSSLValidation bool)
	setTargetInformationMutex       sync.RWMutex
	setTargetInformationArgsForCall []struct {
		api               string
		apiVersion        string
		auth              string
		loggregator       string
		minCLIVersion     string
		doppler           string
		uaa               string
		routing           string
		skipSSLValidation bool
	}
	SetTokenInformationStub        func(accessToken string, refreshToken string, sshOAuthClient string)
	setTokenInformationMutex       sync.RWMutex
	setTokenInformationArgsForCall []struct {
		accessToken    string
		refreshToken   string
		sshOAuthClient string
	}
	SkipSSLValidationStub        func() bool
	skipSSLValidationMutex       sync.RWMutex
	skipSSLValidationArgsForCall []struct{}
	skipSSLValidationReturns     struct {
		result1 bool
	}
	TargetStub        func() string
	targetMutex       sync.RWMutex
	targetArgsForCall []struct{}
	targetReturns     struct {
		result1 string
	}
	TargetedOrganizationStub        func() configv3.Organization
	targetedOrganizationMutex       sync.RWMutex
	targetedOrganizationArgsForCall []struct{}
	targetedOrganizationReturns     struct {
		result1 configv3.Organization
	}
	TargetedSpaceStub        func() configv3.Space
	targetedSpaceMutex       sync.RWMutex
	targetedSpaceArgsForCall []struct{}
	targetedSpaceReturns     struct {
		result1 configv3.Space
	}
	UAAOAuthClientStub        func() string
	uAAOAuthClientMutex       sync.RWMutex
	uAAOAuthClientArgsForCall []struct{}
	uAAOAuthClientReturns     struct {
		result1 string
	}
	UAAOAuthClientSecretStub        func() string
	uAAOAuthClientSecretMutex       sync.RWMutex
	uAAOAuthClientSecretArgsForCall []struct{}
	uAAOAuthClientSecretReturns     struct {
		result1 string
	}
	VerboseStub        func() (bool, []string)
	verboseMutex       sync.RWMutex
	verboseArgsForCall []struct{}
	verboseReturns     struct {
		result1 bool
		result2 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) APIVersion() string {
	fake.aPIVersionMutex.Lock()
	fake.aPIVersionArgsForCall = append(fake.aPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("APIVersion", []interface{}{})
	fake.aPIVersionMutex.Unlock()
	if fake.APIVersionStub != nil {
		return fake.APIVersionStub()
	} else {
		return fake.aPIVersionReturns.result1
	}
}

func (fake *FakeConfig) APIVersionCallCount() int {
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	return len(fake.aPIVersionArgsForCall)
}

func (fake *FakeConfig) APIVersionReturns(result1 string) {
	fake.APIVersionStub = nil
	fake.aPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) AccessToken() string {
	fake.accessTokenMutex.Lock()
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct{}{})
	fake.recordInvocation("AccessToken", []interface{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	} else {
		return fake.accessTokenReturns.result1
	}
}

func (fake *FakeConfig) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeConfig) AccessTokenReturns(result1 string) {
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) BinaryBuildDate() string {
	fake.binaryBuildDateMutex.Lock()
	fake.binaryBuildDateArgsForCall = append(fake.binaryBuildDateArgsForCall, struct{}{})
	fake.recordInvocation("BinaryBuildDate", []interface{}{})
	fake.binaryBuildDateMutex.Unlock()
	if fake.BinaryBuildDateStub != nil {
		return fake.BinaryBuildDateStub()
	} else {
		return fake.binaryBuildDateReturns.result1
	}
}

func (fake *FakeConfig) BinaryBuildDateCallCount() int {
	fake.binaryBuildDateMutex.RLock()
	defer fake.binaryBuildDateMutex.RUnlock()
	return len(fake.binaryBuildDateArgsForCall)
}

func (fake *FakeConfig) BinaryBuildDateReturns(result1 string) {
	fake.BinaryBuildDateStub = nil
	fake.binaryBuildDateReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) BinaryName() string {
	fake.binaryNameMutex.Lock()
	fake.binaryNameArgsForCall = append(fake.binaryNameArgsForCall, struct{}{})
	fake.recordInvocation("BinaryName", []interface{}{})
	fake.binaryNameMutex.Unlock()
	if fake.BinaryNameStub != nil {
		return fake.BinaryNameStub()
	} else {
		return fake.binaryNameReturns.result1
	}
}

func (fake *FakeConfig) BinaryNameCallCount() int {
	fake.binaryNameMutex.RLock()
	defer fake.binaryNameMutex.RUnlock()
	return len(fake.binaryNameArgsForCall)
}

func (fake *FakeConfig) BinaryNameReturns(result1 string) {
	fake.BinaryNameStub = nil
	fake.binaryNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) BinaryVersion() string {
	fake.binaryVersionMutex.Lock()
	fake.binaryVersionArgsForCall = append(fake.binaryVersionArgsForCall, struct{}{})
	fake.recordInvocation("BinaryVersion", []interface{}{})
	fake.binaryVersionMutex.Unlock()
	if fake.BinaryVersionStub != nil {
		return fake.BinaryVersionStub()
	} else {
		return fake.binaryVersionReturns.result1
	}
}

func (fake *FakeConfig) BinaryVersionCallCount() int {
	fake.binaryVersionMutex.RLock()
	defer fake.binaryVersionMutex.RUnlock()
	return len(fake.binaryVersionArgsForCall)
}

func (fake *FakeConfig) BinaryVersionReturns(result1 string) {
	fake.BinaryVersionStub = nil
	fake.binaryVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) ColorEnabled() configv3.ColorSetting {
	fake.colorEnabledMutex.Lock()
	fake.colorEnabledArgsForCall = append(fake.colorEnabledArgsForCall, struct{}{})
	fake.recordInvocation("ColorEnabled", []interface{}{})
	fake.colorEnabledMutex.Unlock()
	if fake.ColorEnabledStub != nil {
		return fake.ColorEnabledStub()
	} else {
		return fake.colorEnabledReturns.result1
	}
}

func (fake *FakeConfig) ColorEnabledCallCount() int {
	fake.colorEnabledMutex.RLock()
	defer fake.colorEnabledMutex.RUnlock()
	return len(fake.colorEnabledArgsForCall)
}

func (fake *FakeConfig) ColorEnabledReturns(result1 configv3.ColorSetting) {
	fake.ColorEnabledStub = nil
	fake.colorEnabledReturns = struct {
		result1 configv3.ColorSetting
	}{result1}
}

func (fake *FakeConfig) CurrentUser() (configv3.User, error) {
	fake.currentUserMutex.Lock()
	fake.currentUserArgsForCall = append(fake.currentUserArgsForCall, struct{}{})
	fake.recordInvocation("CurrentUser", []interface{}{})
	fake.currentUserMutex.Unlock()
	if fake.CurrentUserStub != nil {
		return fake.CurrentUserStub()
	} else {
		return fake.currentUserReturns.result1, fake.currentUserReturns.result2
	}
}

func (fake *FakeConfig) CurrentUserCallCount() int {
	fake.currentUserMutex.RLock()
	defer fake.currentUserMutex.RUnlock()
	return len(fake.currentUserArgsForCall)
}

func (fake *FakeConfig) CurrentUserReturns(result1 configv3.User, result2 error) {
	fake.CurrentUserStub = nil
	fake.currentUserReturns = struct {
		result1 configv3.User
		result2 error
	}{result1, result2}
}

func (fake *FakeConfig) DialTimeout() time.Duration {
	fake.dialTimeoutMutex.Lock()
	fake.dialTimeoutArgsForCall = append(fake.dialTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("DialTimeout", []interface{}{})
	fake.dialTimeoutMutex.Unlock()
	if fake.DialTimeoutStub != nil {
		return fake.DialTimeoutStub()
	} else {
		return fake.dialTimeoutReturns.result1
	}
}

func (fake *FakeConfig) DialTimeoutCallCount() int {
	fake.dialTimeoutMutex.RLock()
	defer fake.dialTimeoutMutex.RUnlock()
	return len(fake.dialTimeoutArgsForCall)
}

func (fake *FakeConfig) DialTimeoutReturns(result1 time.Duration) {
	fake.DialTimeoutStub = nil
	fake.dialTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) Experimental() bool {
	fake.experimentalMutex.Lock()
	fake.experimentalArgsForCall = append(fake.experimentalArgsForCall, struct{}{})
	fake.recordInvocation("Experimental", []interface{}{})
	fake.experimentalMutex.Unlock()
	if fake.ExperimentalStub != nil {
		return fake.ExperimentalStub()
	} else {
		return fake.experimentalReturns.result1
	}
}

func (fake *FakeConfig) ExperimentalCallCount() int {
	fake.experimentalMutex.RLock()
	defer fake.experimentalMutex.RUnlock()
	return len(fake.experimentalArgsForCall)
}

func (fake *FakeConfig) ExperimentalReturns(result1 bool) {
	fake.ExperimentalStub = nil
	fake.experimentalReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) Locale() string {
	fake.localeMutex.Lock()
	fake.localeArgsForCall = append(fake.localeArgsForCall, struct{}{})
	fake.recordInvocation("Locale", []interface{}{})
	fake.localeMutex.Unlock()
	if fake.LocaleStub != nil {
		return fake.LocaleStub()
	} else {
		return fake.localeReturns.result1
	}
}

func (fake *FakeConfig) LocaleCallCount() int {
	fake.localeMutex.RLock()
	defer fake.localeMutex.RUnlock()
	return len(fake.localeArgsForCall)
}

func (fake *FakeConfig) LocaleReturns(result1 string) {
	fake.LocaleStub = nil
	fake.localeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) MinCLIVersion() string {
	fake.minCLIVersionMutex.Lock()
	fake.minCLIVersionArgsForCall = append(fake.minCLIVersionArgsForCall, struct{}{})
	fake.recordInvocation("MinCLIVersion", []interface{}{})
	fake.minCLIVersionMutex.Unlock()
	if fake.MinCLIVersionStub != nil {
		return fake.MinCLIVersionStub()
	} else {
		return fake.minCLIVersionReturns.result1
	}
}

func (fake *FakeConfig) MinCLIVersionCallCount() int {
	fake.minCLIVersionMutex.RLock()
	defer fake.minCLIVersionMutex.RUnlock()
	return len(fake.minCLIVersionArgsForCall)
}

func (fake *FakeConfig) MinCLIVersionReturns(result1 string) {
	fake.MinCLIVersionStub = nil
	fake.minCLIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) Plugins() map[string]configv3.Plugin {
	fake.pluginsMutex.Lock()
	fake.pluginsArgsForCall = append(fake.pluginsArgsForCall, struct{}{})
	fake.recordInvocation("Plugins", []interface{}{})
	fake.pluginsMutex.Unlock()
	if fake.PluginsStub != nil {
		return fake.PluginsStub()
	} else {
		return fake.pluginsReturns.result1
	}
}

func (fake *FakeConfig) PluginsCallCount() int {
	fake.pluginsMutex.RLock()
	defer fake.pluginsMutex.RUnlock()
	return len(fake.pluginsArgsForCall)
}

func (fake *FakeConfig) PluginsReturns(result1 map[string]configv3.Plugin) {
	fake.PluginsStub = nil
	fake.pluginsReturns = struct {
		result1 map[string]configv3.Plugin
	}{result1}
}

func (fake *FakeConfig) RefreshToken() string {
	fake.refreshTokenMutex.Lock()
	fake.refreshTokenArgsForCall = append(fake.refreshTokenArgsForCall, struct{}{})
	fake.recordInvocation("RefreshToken", []interface{}{})
	fake.refreshTokenMutex.Unlock()
	if fake.RefreshTokenStub != nil {
		return fake.RefreshTokenStub()
	} else {
		return fake.refreshTokenReturns.result1
	}
}

func (fake *FakeConfig) RefreshTokenCallCount() int {
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	return len(fake.refreshTokenArgsForCall)
}

func (fake *FakeConfig) RefreshTokenReturns(result1 string) {
	fake.RefreshTokenStub = nil
	fake.refreshTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) PollingInterval() time.Duration {
	fake.pollingIntervalMutex.Lock()
	fake.pollingIntervalArgsForCall = append(fake.pollingIntervalArgsForCall, struct{}{})
	fake.recordInvocation("PollingInterval", []interface{}{})
	fake.pollingIntervalMutex.Unlock()
	if fake.PollingIntervalStub != nil {
		return fake.PollingIntervalStub()
	} else {
		return fake.pollingIntervalReturns.result1
	}
}

func (fake *FakeConfig) PollingIntervalCallCount() int {
	fake.pollingIntervalMutex.RLock()
	defer fake.pollingIntervalMutex.RUnlock()
	return len(fake.pollingIntervalArgsForCall)
}

func (fake *FakeConfig) PollingIntervalReturns(result1 time.Duration) {
	fake.PollingIntervalStub = nil
	fake.pollingIntervalReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) OverallPollingTimeout() time.Duration {
	fake.overallPollingTimeoutMutex.Lock()
	fake.overallPollingTimeoutArgsForCall = append(fake.overallPollingTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("OverallPollingTimeout", []interface{}{})
	fake.overallPollingTimeoutMutex.Unlock()
	if fake.OverallPollingTimeoutStub != nil {
		return fake.OverallPollingTimeoutStub()
	} else {
		return fake.overallPollingTimeoutReturns.result1
	}
}

func (fake *FakeConfig) OverallPollingTimeoutCallCount() int {
	fake.overallPollingTimeoutMutex.RLock()
	defer fake.overallPollingTimeoutMutex.RUnlock()
	return len(fake.overallPollingTimeoutArgsForCall)
}

func (fake *FakeConfig) OverallPollingTimeoutReturns(result1 time.Duration) {
	fake.OverallPollingTimeoutStub = nil
	fake.overallPollingTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) SetAccessToken(token string) {
	fake.setAccessTokenMutex.Lock()
	fake.setAccessTokenArgsForCall = append(fake.setAccessTokenArgsForCall, struct {
		token string
	}{token})
	fake.recordInvocation("SetAccessToken", []interface{}{token})
	fake.setAccessTokenMutex.Unlock()
	if fake.SetAccessTokenStub != nil {
		fake.SetAccessTokenStub(token)
	}
}

func (fake *FakeConfig) SetAccessTokenCallCount() int {
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	return len(fake.setAccessTokenArgsForCall)
}

func (fake *FakeConfig) SetAccessTokenArgsForCall(i int) string {
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	return fake.setAccessTokenArgsForCall[i].token
}

func (fake *FakeConfig) SetOrganizationInformation(guid string, name string) {
	fake.setOrganizationInformationMutex.Lock()
	fake.setOrganizationInformationArgsForCall = append(fake.setOrganizationInformationArgsForCall, struct {
		guid string
		name string
	}{guid, name})
	fake.recordInvocation("SetOrganizationInformation", []interface{}{guid, name})
	fake.setOrganizationInformationMutex.Unlock()
	if fake.SetOrganizationInformationStub != nil {
		fake.SetOrganizationInformationStub(guid, name)
	}
}

func (fake *FakeConfig) SetOrganizationInformationCallCount() int {
	fake.setOrganizationInformationMutex.RLock()
	defer fake.setOrganizationInformationMutex.RUnlock()
	return len(fake.setOrganizationInformationArgsForCall)
}

func (fake *FakeConfig) SetOrganizationInformationArgsForCall(i int) (string, string) {
	fake.setOrganizationInformationMutex.RLock()
	defer fake.setOrganizationInformationMutex.RUnlock()
	return fake.setOrganizationInformationArgsForCall[i].guid, fake.setOrganizationInformationArgsForCall[i].name
}

func (fake *FakeConfig) SetSpaceInformation(guid string, name string, allowSSH bool) {
	fake.setSpaceInformationMutex.Lock()
	fake.setSpaceInformationArgsForCall = append(fake.setSpaceInformationArgsForCall, struct {
		guid     string
		name     string
		allowSSH bool
	}{guid, name, allowSSH})
	fake.recordInvocation("SetSpaceInformation", []interface{}{guid, name, allowSSH})
	fake.setSpaceInformationMutex.Unlock()
	if fake.SetSpaceInformationStub != nil {
		fake.SetSpaceInformationStub(guid, name, allowSSH)
	}
}

func (fake *FakeConfig) SetSpaceInformationCallCount() int {
	fake.setSpaceInformationMutex.RLock()
	defer fake.setSpaceInformationMutex.RUnlock()
	return len(fake.setSpaceInformationArgsForCall)
}

func (fake *FakeConfig) SetSpaceInformationArgsForCall(i int) (string, string, bool) {
	fake.setSpaceInformationMutex.RLock()
	defer fake.setSpaceInformationMutex.RUnlock()
	return fake.setSpaceInformationArgsForCall[i].guid, fake.setSpaceInformationArgsForCall[i].name, fake.setSpaceInformationArgsForCall[i].allowSSH
}

func (fake *FakeConfig) UnsetSpaceInformation() {
	fake.unsetSpaceInformationMutex.Lock()
	fake.unsetSpaceInformationArgsForCall = append(fake.unsetSpaceInformationArgsForCall, struct{}{})
	fake.recordInvocation("UnsetSpaceInformation", []interface{}{})
	fake.unsetSpaceInformationMutex.Unlock()
	if fake.UnsetSpaceInformationStub != nil {
		fake.UnsetSpaceInformationStub()
	}
}

func (fake *FakeConfig) UnsetSpaceInformationCallCount() int {
	fake.unsetSpaceInformationMutex.RLock()
	defer fake.unsetSpaceInformationMutex.RUnlock()
	return len(fake.unsetSpaceInformationArgsForCall)
}

func (fake *FakeConfig) SetRefreshToken(token string) {
	fake.setRefreshTokenMutex.Lock()
	fake.setRefreshTokenArgsForCall = append(fake.setRefreshTokenArgsForCall, struct {
		token string
	}{token})
	fake.recordInvocation("SetRefreshToken", []interface{}{token})
	fake.setRefreshTokenMutex.Unlock()
	if fake.SetRefreshTokenStub != nil {
		fake.SetRefreshTokenStub(token)
	}
}

func (fake *FakeConfig) SetRefreshTokenCallCount() int {
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	return len(fake.setRefreshTokenArgsForCall)
}

func (fake *FakeConfig) SetRefreshTokenArgsForCall(i int) string {
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	return fake.setRefreshTokenArgsForCall[i].token
}

func (fake *FakeConfig) SetTargetInformation(api string, apiVersion string, auth string, loggregator string, minCLIVersion string, doppler string, uaa string, routing string, skipSSLValidation bool) {
	fake.setTargetInformationMutex.Lock()
	fake.setTargetInformationArgsForCall = append(fake.setTargetInformationArgsForCall, struct {
		api               string
		apiVersion        string
		auth              string
		loggregator       string
		minCLIVersion     string
		doppler           string
		uaa               string
		routing           string
		skipSSLValidation bool
	}{api, apiVersion, auth, loggregator, minCLIVersion, doppler, uaa, routing, skipSSLValidation})
	fake.recordInvocation("SetTargetInformation", []interface{}{api, apiVersion, auth, loggregator, minCLIVersion, doppler, uaa, routing, skipSSLValidation})
	fake.setTargetInformationMutex.Unlock()
	if fake.SetTargetInformationStub != nil {
		fake.SetTargetInformationStub(api, apiVersion, auth, loggregator, minCLIVersion, doppler, uaa, routing, skipSSLValidation)
	}
}

func (fake *FakeConfig) SetTargetInformationCallCount() int {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return len(fake.setTargetInformationArgsForCall)
}

func (fake *FakeConfig) SetTargetInformationArgsForCall(i int) (string, string, string, string, string, string, string, string, bool) {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return fake.setTargetInformationArgsForCall[i].api, fake.setTargetInformationArgsForCall[i].apiVersion, fake.setTargetInformationArgsForCall[i].auth, fake.setTargetInformationArgsForCall[i].loggregator, fake.setTargetInformationArgsForCall[i].minCLIVersion, fake.setTargetInformationArgsForCall[i].doppler, fake.setTargetInformationArgsForCall[i].uaa, fake.setTargetInformationArgsForCall[i].routing, fake.setTargetInformationArgsForCall[i].skipSSLValidation
}

func (fake *FakeConfig) SetTokenInformation(accessToken string, refreshToken string, sshOAuthClient string) {
	fake.setTokenInformationMutex.Lock()
	fake.setTokenInformationArgsForCall = append(fake.setTokenInformationArgsForCall, struct {
		accessToken    string
		refreshToken   string
		sshOAuthClient string
	}{accessToken, refreshToken, sshOAuthClient})
	fake.recordInvocation("SetTokenInformation", []interface{}{accessToken, refreshToken, sshOAuthClient})
	fake.setTokenInformationMutex.Unlock()
	if fake.SetTokenInformationStub != nil {
		fake.SetTokenInformationStub(accessToken, refreshToken, sshOAuthClient)
	}
}

func (fake *FakeConfig) SetTokenInformationCallCount() int {
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	return len(fake.setTokenInformationArgsForCall)
}

func (fake *FakeConfig) SetTokenInformationArgsForCall(i int) (string, string, string) {
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	return fake.setTokenInformationArgsForCall[i].accessToken, fake.setTokenInformationArgsForCall[i].refreshToken, fake.setTokenInformationArgsForCall[i].sshOAuthClient
}

func (fake *FakeConfig) SkipSSLValidation() bool {
	fake.skipSSLValidationMutex.Lock()
	fake.skipSSLValidationArgsForCall = append(fake.skipSSLValidationArgsForCall, struct{}{})
	fake.recordInvocation("SkipSSLValidation", []interface{}{})
	fake.skipSSLValidationMutex.Unlock()
	if fake.SkipSSLValidationStub != nil {
		return fake.SkipSSLValidationStub()
	} else {
		return fake.skipSSLValidationReturns.result1
	}
}

func (fake *FakeConfig) SkipSSLValidationCallCount() int {
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	return len(fake.skipSSLValidationArgsForCall)
}

func (fake *FakeConfig) SkipSSLValidationReturns(result1 bool) {
	fake.SkipSSLValidationStub = nil
	fake.skipSSLValidationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) Target() string {
	fake.targetMutex.Lock()
	fake.targetArgsForCall = append(fake.targetArgsForCall, struct{}{})
	fake.recordInvocation("Target", []interface{}{})
	fake.targetMutex.Unlock()
	if fake.TargetStub != nil {
		return fake.TargetStub()
	} else {
		return fake.targetReturns.result1
	}
}

func (fake *FakeConfig) TargetCallCount() int {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return len(fake.targetArgsForCall)
}

func (fake *FakeConfig) TargetReturns(result1 string) {
	fake.TargetStub = nil
	fake.targetReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) TargetedOrganization() configv3.Organization {
	fake.targetedOrganizationMutex.Lock()
	fake.targetedOrganizationArgsForCall = append(fake.targetedOrganizationArgsForCall, struct{}{})
	fake.recordInvocation("TargetedOrganization", []interface{}{})
	fake.targetedOrganizationMutex.Unlock()
	if fake.TargetedOrganizationStub != nil {
		return fake.TargetedOrganizationStub()
	} else {
		return fake.targetedOrganizationReturns.result1
	}
}

func (fake *FakeConfig) TargetedOrganizationCallCount() int {
	fake.targetedOrganizationMutex.RLock()
	defer fake.targetedOrganizationMutex.RUnlock()
	return len(fake.targetedOrganizationArgsForCall)
}

func (fake *FakeConfig) TargetedOrganizationReturns(result1 configv3.Organization) {
	fake.TargetedOrganizationStub = nil
	fake.targetedOrganizationReturns = struct {
		result1 configv3.Organization
	}{result1}
}

func (fake *FakeConfig) TargetedSpace() configv3.Space {
	fake.targetedSpaceMutex.Lock()
	fake.targetedSpaceArgsForCall = append(fake.targetedSpaceArgsForCall, struct{}{})
	fake.recordInvocation("TargetedSpace", []interface{}{})
	fake.targetedSpaceMutex.Unlock()
	if fake.TargetedSpaceStub != nil {
		return fake.TargetedSpaceStub()
	} else {
		return fake.targetedSpaceReturns.result1
	}
}

func (fake *FakeConfig) TargetedSpaceCallCount() int {
	fake.targetedSpaceMutex.RLock()
	defer fake.targetedSpaceMutex.RUnlock()
	return len(fake.targetedSpaceArgsForCall)
}

func (fake *FakeConfig) TargetedSpaceReturns(result1 configv3.Space) {
	fake.TargetedSpaceStub = nil
	fake.targetedSpaceReturns = struct {
		result1 configv3.Space
	}{result1}
}

func (fake *FakeConfig) UAAOAuthClient() string {
	fake.uAAOAuthClientMutex.Lock()
	fake.uAAOAuthClientArgsForCall = append(fake.uAAOAuthClientArgsForCall, struct{}{})
	fake.recordInvocation("UAAOAuthClient", []interface{}{})
	fake.uAAOAuthClientMutex.Unlock()
	if fake.UAAOAuthClientStub != nil {
		return fake.UAAOAuthClientStub()
	} else {
		return fake.uAAOAuthClientReturns.result1
	}
}

func (fake *FakeConfig) UAAOAuthClientCallCount() int {
	fake.uAAOAuthClientMutex.RLock()
	defer fake.uAAOAuthClientMutex.RUnlock()
	return len(fake.uAAOAuthClientArgsForCall)
}

func (fake *FakeConfig) UAAOAuthClientReturns(result1 string) {
	fake.UAAOAuthClientStub = nil
	fake.uAAOAuthClientReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) UAAOAuthClientSecret() string {
	fake.uAAOAuthClientSecretMutex.Lock()
	fake.uAAOAuthClientSecretArgsForCall = append(fake.uAAOAuthClientSecretArgsForCall, struct{}{})
	fake.recordInvocation("UAAOAuthClientSecret", []interface{}{})
	fake.uAAOAuthClientSecretMutex.Unlock()
	if fake.UAAOAuthClientSecretStub != nil {
		return fake.UAAOAuthClientSecretStub()
	} else {
		return fake.uAAOAuthClientSecretReturns.result1
	}
}

func (fake *FakeConfig) UAAOAuthClientSecretCallCount() int {
	fake.uAAOAuthClientSecretMutex.RLock()
	defer fake.uAAOAuthClientSecretMutex.RUnlock()
	return len(fake.uAAOAuthClientSecretArgsForCall)
}

func (fake *FakeConfig) UAAOAuthClientSecretReturns(result1 string) {
	fake.UAAOAuthClientSecretStub = nil
	fake.uAAOAuthClientSecretReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) Verbose() (bool, []string) {
	fake.verboseMutex.Lock()
	fake.verboseArgsForCall = append(fake.verboseArgsForCall, struct{}{})
	fake.recordInvocation("Verbose", []interface{}{})
	fake.verboseMutex.Unlock()
	if fake.VerboseStub != nil {
		return fake.VerboseStub()
	} else {
		return fake.verboseReturns.result1, fake.verboseReturns.result2
	}
}

func (fake *FakeConfig) VerboseCallCount() int {
	fake.verboseMutex.RLock()
	defer fake.verboseMutex.RUnlock()
	return len(fake.verboseArgsForCall)
}

func (fake *FakeConfig) VerboseReturns(result1 bool, result2 []string) {
	fake.VerboseStub = nil
	fake.verboseReturns = struct {
		result1 bool
		result2 []string
	}{result1, result2}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	fake.binaryBuildDateMutex.RLock()
	defer fake.binaryBuildDateMutex.RUnlock()
	fake.binaryNameMutex.RLock()
	defer fake.binaryNameMutex.RUnlock()
	fake.binaryVersionMutex.RLock()
	defer fake.binaryVersionMutex.RUnlock()
	fake.colorEnabledMutex.RLock()
	defer fake.colorEnabledMutex.RUnlock()
	fake.currentUserMutex.RLock()
	defer fake.currentUserMutex.RUnlock()
	fake.dialTimeoutMutex.RLock()
	defer fake.dialTimeoutMutex.RUnlock()
	fake.experimentalMutex.RLock()
	defer fake.experimentalMutex.RUnlock()
	fake.localeMutex.RLock()
	defer fake.localeMutex.RUnlock()
	fake.minCLIVersionMutex.RLock()
	defer fake.minCLIVersionMutex.RUnlock()
	fake.pluginsMutex.RLock()
	defer fake.pluginsMutex.RUnlock()
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	fake.pollingIntervalMutex.RLock()
	defer fake.pollingIntervalMutex.RUnlock()
	fake.overallPollingTimeoutMutex.RLock()
	defer fake.overallPollingTimeoutMutex.RUnlock()
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	fake.setOrganizationInformationMutex.RLock()
	defer fake.setOrganizationInformationMutex.RUnlock()
	fake.setSpaceInformationMutex.RLock()
	defer fake.setSpaceInformationMutex.RUnlock()
	fake.unsetSpaceInformationMutex.RLock()
	defer fake.unsetSpaceInformationMutex.RUnlock()
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	fake.targetedOrganizationMutex.RLock()
	defer fake.targetedOrganizationMutex.RUnlock()
	fake.targetedSpaceMutex.RLock()
	defer fake.targetedSpaceMutex.RUnlock()
	fake.uAAOAuthClientMutex.RLock()
	defer fake.uAAOAuthClientMutex.RUnlock()
	fake.uAAOAuthClientSecretMutex.RLock()
	defer fake.uAAOAuthClientSecretMutex.RUnlock()
	fake.verboseMutex.RLock()
	defer fake.verboseMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ command.Config = new(FakeConfig)
