<?php
/**
 * SelfServiceRecoveryFlowState
 *
 * PHP version 7.3
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.
 *
 * The version of the OpenAPI document: v0.0.1-alpha.18
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.2.1
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace Ory\Client\Model;
use \Ory\Client\ObjectSerializer;

/**
 * SelfServiceRecoveryFlowState Class Doc Comment
 *
 * @category Class
 * @description The state represents the state of the recovery flow.  choose_method: ask the user to choose a method (e.g. recover account via email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the recovery challenge was passed.
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class SelfServiceRecoveryFlowState
{
    /**
     * Possible values of this enum
     */
    const CHOOSE_METHOD = 'choose_method';
    const SENT_EMAIL = 'sent_email';
    const PASSED_CHALLENGE = 'passed_challenge';
    
    /**
     * Gets allowable values of the enum
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [
            self::CHOOSE_METHOD,
            self::SENT_EMAIL,
            self::PASSED_CHALLENGE,
        ];
    }
}


