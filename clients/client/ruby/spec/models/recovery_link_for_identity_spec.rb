=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v1.0.0
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OryClient::RecoveryLinkForIdentity
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OryClient::RecoveryLinkForIdentity do
  let(:instance) { OryClient::RecoveryLinkForIdentity.new }

  describe 'test an instance of RecoveryLinkForIdentity' do
    it 'should create an instance of RecoveryLinkForIdentity' do
      expect(instance).to be_instance_of(OryClient::RecoveryLinkForIdentity)
    end
  end
  describe 'test attribute "expires_at"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "recovery_link"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end