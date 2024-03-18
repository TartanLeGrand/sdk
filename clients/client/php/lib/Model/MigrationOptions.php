<?php
/**
 * MigrationOptions
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
 * The version of the OpenAPI document: v1.8.1
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.4.0
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
 * MigrationOptions Class Doc Comment
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class MigrationOptions implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'migrationOptions';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'environment' => 'string',
        'projectSubscription' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'environment' => null,
        'projectSubscription' => null
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
        'environment' => 'environment',
        'projectSubscription' => 'project_subscription'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'environment' => 'setEnvironment',
        'projectSubscription' => 'setProjectSubscription'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'environment' => 'getEnvironment',
        'projectSubscription' => 'getProjectSubscription'
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

    const ENVIRONMENT_PROD = 'prod';
    const ENVIRONMENT_DEV = 'dev';
    const PROJECT_SUBSCRIPTION_MIGRATE = 'migrate';
    const PROJECT_SUBSCRIPTION_IGNORE = 'ignore';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getEnvironmentAllowableValues()
    {
        return [
            self::ENVIRONMENT_PROD,
            self::ENVIRONMENT_DEV,
        ];
    }

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getProjectSubscriptionAllowableValues()
    {
        return [
            self::PROJECT_SUBSCRIPTION_MIGRATE,
            self::PROJECT_SUBSCRIPTION_IGNORE,
        ];
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
        $this->container['environment'] = $data['environment'] ?? null;
        $this->container['projectSubscription'] = $data['projectSubscription'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['environment'] === null) {
            $invalidProperties[] = "'environment' can't be null";
        }
        $allowedValues = $this->getEnvironmentAllowableValues();
        if (!is_null($this->container['environment']) && !in_array($this->container['environment'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'environment', must be one of '%s'",
                $this->container['environment'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['projectSubscription'] === null) {
            $invalidProperties[] = "'projectSubscription' can't be null";
        }
        $allowedValues = $this->getProjectSubscriptionAllowableValues();
        if (!is_null($this->container['projectSubscription']) && !in_array($this->container['projectSubscription'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'projectSubscription', must be one of '%s'",
                $this->container['projectSubscription'],
                implode("', '", $allowedValues)
            );
        }

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
     * Gets environment
     *
     * @return string
     */
    public function getEnvironment()
    {
        return $this->container['environment'];
    }

    /**
     * Sets environment
     *
     * @param string $environment The environment of the project in the workspace. Can be one of \"prod\" or \"dev\". Note that the number of projects in the \"prod\" environment is limited depending on the subscription. prod Production dev Development
     *
     * @return self
     */
    public function setEnvironment($environment)
    {
        $allowedValues = $this->getEnvironmentAllowableValues();
        if (!in_array($environment, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'environment', must be one of '%s'",
                    $environment,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['environment'] = $environment;

        return $this;
    }

    /**
     * Gets projectSubscription
     *
     * @return string
     */
    public function getProjectSubscription()
    {
        return $this->container['projectSubscription'];
    }

    /**
     * Sets projectSubscription
     *
     * @param string $projectSubscription The action to take with the project subscription. Can be one of \"migrate\", and \"ignore\". \"migrate\" will migrate the project subscription to the workspace. \"ignore\" will ignore the project subscription. migrate ProjectSubscriptionActionMigrate  ProjectSubscriptionActionMigrate will migrate the project subscription to the  workspace. ignore ProjectSubscriptionActionIgnore  ProjectSubscriptionActionIgnore will ignore the project subscription.
     *
     * @return self
     */
    public function setProjectSubscription($projectSubscription)
    {
        $allowedValues = $this->getProjectSubscriptionAllowableValues();
        if (!in_array($projectSubscription, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'projectSubscription', must be one of '%s'",
                    $projectSubscription,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['projectSubscription'] = $projectSubscription;

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

