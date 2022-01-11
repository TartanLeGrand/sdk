<?php
/**
 * RedirectionConfig
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
 * The version of the OpenAPI document: v0.0.1-alpha.44
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

use \ArrayAccess;
use \Ory\Client\ObjectSerializer;

/**
 * RedirectionConfig Class Doc Comment
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class RedirectionConfig implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'RedirectionConfig';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'global' => '\Ory\Client\Model\RedirectionField',
        'login' => '\Ory\Client\Model\RedirectionField',
        'logout' => '\Ory\Client\Model\RedirectionField',
        'recovery' => '\Ory\Client\Model\RedirectionField',
        'registration' => '\Ory\Client\Model\RedirectionField',
        'settings' => '\Ory\Client\Model\RedirectionField',
        'urlAllowlist' => 'string[]',
        'verification' => '\Ory\Client\Model\RedirectionField'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'global' => null,
        'login' => null,
        'logout' => null,
        'recovery' => null,
        'registration' => null,
        'settings' => null,
        'urlAllowlist' => null,
        'verification' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'global' => 'global',
        'login' => 'login',
        'logout' => 'logout',
        'recovery' => 'recovery',
        'registration' => 'registration',
        'settings' => 'settings',
        'urlAllowlist' => 'url_allowlist',
        'verification' => 'verification'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'global' => 'setGlobal',
        'login' => 'setLogin',
        'logout' => 'setLogout',
        'recovery' => 'setRecovery',
        'registration' => 'setRegistration',
        'settings' => 'setSettings',
        'urlAllowlist' => 'setUrlAllowlist',
        'verification' => 'setVerification'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'global' => 'getGlobal',
        'login' => 'getLogin',
        'logout' => 'getLogout',
        'recovery' => 'getRecovery',
        'registration' => 'getRegistration',
        'settings' => 'getSettings',
        'urlAllowlist' => 'getUrlAllowlist',
        'verification' => 'getVerification'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }


    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['global'] = $data['global'] ?? null;
        $this->container['login'] = $data['login'] ?? null;
        $this->container['logout'] = $data['logout'] ?? null;
        $this->container['recovery'] = $data['recovery'] ?? null;
        $this->container['registration'] = $data['registration'] ?? null;
        $this->container['settings'] = $data['settings'] ?? null;
        $this->container['urlAllowlist'] = $data['urlAllowlist'] ?? null;
        $this->container['verification'] = $data['verification'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets global
     *
     * @return \Ory\Client\Model\RedirectionField|null
     */
    public function getGlobal()
    {
        return $this->container['global'];
    }

    /**
     * Sets global
     *
     * @param \Ory\Client\Model\RedirectionField|null $global global
     *
     * @return self
     */
    public function setGlobal($global)
    {
        $this->container['global'] = $global;

        return $this;
    }

    /**
     * Gets login
     *
     * @return \Ory\Client\Model\RedirectionField|null
     */
    public function getLogin()
    {
        return $this->container['login'];
    }

    /**
     * Sets login
     *
     * @param \Ory\Client\Model\RedirectionField|null $login login
     *
     * @return self
     */
    public function setLogin($login)
    {
        $this->container['login'] = $login;

        return $this;
    }

    /**
     * Gets logout
     *
     * @return \Ory\Client\Model\RedirectionField|null
     */
    public function getLogout()
    {
        return $this->container['logout'];
    }

    /**
     * Sets logout
     *
     * @param \Ory\Client\Model\RedirectionField|null $logout logout
     *
     * @return self
     */
    public function setLogout($logout)
    {
        $this->container['logout'] = $logout;

        return $this;
    }

    /**
     * Gets recovery
     *
     * @return \Ory\Client\Model\RedirectionField|null
     */
    public function getRecovery()
    {
        return $this->container['recovery'];
    }

    /**
     * Sets recovery
     *
     * @param \Ory\Client\Model\RedirectionField|null $recovery recovery
     *
     * @return self
     */
    public function setRecovery($recovery)
    {
        $this->container['recovery'] = $recovery;

        return $this;
    }

    /**
     * Gets registration
     *
     * @return \Ory\Client\Model\RedirectionField|null
     */
    public function getRegistration()
    {
        return $this->container['registration'];
    }

    /**
     * Sets registration
     *
     * @param \Ory\Client\Model\RedirectionField|null $registration registration
     *
     * @return self
     */
    public function setRegistration($registration)
    {
        $this->container['registration'] = $registration;

        return $this;
    }

    /**
     * Gets settings
     *
     * @return \Ory\Client\Model\RedirectionField|null
     */
    public function getSettings()
    {
        return $this->container['settings'];
    }

    /**
     * Sets settings
     *
     * @param \Ory\Client\Model\RedirectionField|null $settings settings
     *
     * @return self
     */
    public function setSettings($settings)
    {
        $this->container['settings'] = $settings;

        return $this;
    }

    /**
     * Gets urlAllowlist
     *
     * @return string[]|null
     */
    public function getUrlAllowlist()
    {
        return $this->container['urlAllowlist'];
    }

    /**
     * Sets urlAllowlist
     *
     * @param string[]|null $urlAllowlist urlAllowlist
     *
     * @return self
     */
    public function setUrlAllowlist($urlAllowlist)
    {
        $this->container['urlAllowlist'] = $urlAllowlist;

        return $this;
    }

    /**
     * Gets verification
     *
     * @return \Ory\Client\Model\RedirectionField|null
     */
    public function getVerification()
    {
        return $this->container['verification'];
    }

    /**
     * Sets verification
     *
     * @param \Ory\Client\Model\RedirectionField|null $verification verification
     *
     * @return self
     */
    public function setVerification($verification)
    {
        $this->container['verification'] = $verification;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


