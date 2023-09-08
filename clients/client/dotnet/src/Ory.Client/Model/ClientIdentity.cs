/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.4
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Client.Client.OpenAPIDateConverter;

namespace Ory.Client.Model
{
    /// <summary>
    /// An [identity](https://www.ory.sh/docs/kratos/concepts/identity-user-model) represents a (human) user in Ory.
    /// </summary>
    [DataContract(Name = "identity")]
    public partial class ClientIdentity : IEquatable<ClientIdentity>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientIdentity" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientIdentity()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientIdentity" /> class.
        /// </summary>
        /// <param name="createdAt">CreatedAt is a helper struct field for gobuffalo.pop..</param>
        /// <param name="credentials">Credentials represents all credentials that can be used for authenticating this identity..</param>
        /// <param name="id">ID is the identity&#39;s unique identifier.  The Identity ID can not be changed and can not be chosen. This ensures future compatibility and optimization for distributed stores such as CockroachDB. (required).</param>
        /// <param name="metadataAdmin">NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-.</param>
        /// <param name="metadataPublic">NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-.</param>
        /// <param name="recoveryAddresses">RecoveryAddresses contains all the addresses that can be used to recover an identity..</param>
        /// <param name="schemaId">SchemaID is the ID of the JSON Schema to be used for validating the identity&#39;s traits. (required).</param>
        /// <param name="schemaUrl">SchemaURL is the URL of the endpoint where the identity&#39;s traits schema can be fetched from.  format: url (required).</param>
        /// <param name="state">state.</param>
        /// <param name="stateChangedAt">stateChangedAt.</param>
        /// <param name="traits">Traits represent an identity&#39;s traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in &#x60;schema_url&#x60;. (required).</param>
        /// <param name="updatedAt">UpdatedAt is a helper struct field for gobuffalo.pop..</param>
        /// <param name="verifiableAddresses">VerifiableAddresses contains all the addresses that can be verified by the user..</param>
        public ClientIdentity(DateTime createdAt = default(DateTime), Dictionary<string, ClientIdentityCredentials> credentials = default(Dictionary<string, ClientIdentityCredentials>), string id = default(string), Object metadataAdmin = default(Object), Object metadataPublic = default(Object), List<ClientRecoveryIdentityAddress> recoveryAddresses = default(List<ClientRecoveryIdentityAddress>), string schemaId = default(string), string schemaUrl = default(string), ClientIdentityState state = default(ClientIdentityState), DateTime stateChangedAt = default(DateTime), Object traits = default(Object), DateTime updatedAt = default(DateTime), List<ClientVerifiableIdentityAddress> verifiableAddresses = default(List<ClientVerifiableIdentityAddress>))
        {
            // to ensure "id" is required (not null)
            if (id == null) {
                throw new ArgumentNullException("id is a required property for ClientIdentity and cannot be null");
            }
            this.Id = id;
            // to ensure "schemaId" is required (not null)
            if (schemaId == null) {
                throw new ArgumentNullException("schemaId is a required property for ClientIdentity and cannot be null");
            }
            this.SchemaId = schemaId;
            // to ensure "schemaUrl" is required (not null)
            if (schemaUrl == null) {
                throw new ArgumentNullException("schemaUrl is a required property for ClientIdentity and cannot be null");
            }
            this.SchemaUrl = schemaUrl;
            // to ensure "traits" is required (not null)
            if (traits == null) {
                throw new ArgumentNullException("traits is a required property for ClientIdentity and cannot be null");
            }
            this.Traits = traits;
            this.CreatedAt = createdAt;
            this.Credentials = credentials;
            this.MetadataAdmin = metadataAdmin;
            this.MetadataPublic = metadataPublic;
            this.RecoveryAddresses = recoveryAddresses;
            this.State = state;
            this.StateChangedAt = stateChangedAt;
            this.UpdatedAt = updatedAt;
            this.VerifiableAddresses = verifiableAddresses;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// CreatedAt is a helper struct field for gobuffalo.pop.
        /// </summary>
        /// <value>CreatedAt is a helper struct field for gobuffalo.pop.</value>
        [DataMember(Name = "created_at", EmitDefaultValue = false)]
        public DateTime CreatedAt { get; set; }

        /// <summary>
        /// Credentials represents all credentials that can be used for authenticating this identity.
        /// </summary>
        /// <value>Credentials represents all credentials that can be used for authenticating this identity.</value>
        [DataMember(Name = "credentials", EmitDefaultValue = false)]
        public Dictionary<string, ClientIdentityCredentials> Credentials { get; set; }

        /// <summary>
        /// ID is the identity&#39;s unique identifier.  The Identity ID can not be changed and can not be chosen. This ensures future compatibility and optimization for distributed stores such as CockroachDB.
        /// </summary>
        /// <value>ID is the identity&#39;s unique identifier.  The Identity ID can not be changed and can not be chosen. This ensures future compatibility and optimization for distributed stores such as CockroachDB.</value>
        [DataMember(Name = "id", IsRequired = true, EmitDefaultValue = false)]
        public string Id { get; set; }

        /// <summary>
        /// NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-
        /// </summary>
        /// <value>NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-</value>
        [DataMember(Name = "metadata_admin", EmitDefaultValue = true)]
        public Object MetadataAdmin { get; set; }

        /// <summary>
        /// NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-
        /// </summary>
        /// <value>NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-</value>
        [DataMember(Name = "metadata_public", EmitDefaultValue = true)]
        public Object MetadataPublic { get; set; }

        /// <summary>
        /// RecoveryAddresses contains all the addresses that can be used to recover an identity.
        /// </summary>
        /// <value>RecoveryAddresses contains all the addresses that can be used to recover an identity.</value>
        [DataMember(Name = "recovery_addresses", EmitDefaultValue = false)]
        public List<ClientRecoveryIdentityAddress> RecoveryAddresses { get; set; }

        /// <summary>
        /// SchemaID is the ID of the JSON Schema to be used for validating the identity&#39;s traits.
        /// </summary>
        /// <value>SchemaID is the ID of the JSON Schema to be used for validating the identity&#39;s traits.</value>
        [DataMember(Name = "schema_id", IsRequired = true, EmitDefaultValue = false)]
        public string SchemaId { get; set; }

        /// <summary>
        /// SchemaURL is the URL of the endpoint where the identity&#39;s traits schema can be fetched from.  format: url
        /// </summary>
        /// <value>SchemaURL is the URL of the endpoint where the identity&#39;s traits schema can be fetched from.  format: url</value>
        [DataMember(Name = "schema_url", IsRequired = true, EmitDefaultValue = false)]
        public string SchemaUrl { get; set; }

        /// <summary>
        /// Gets or Sets State
        /// </summary>
        [DataMember(Name = "state", EmitDefaultValue = false)]
        public ClientIdentityState State { get; set; }

        /// <summary>
        /// Gets or Sets StateChangedAt
        /// </summary>
        [DataMember(Name = "state_changed_at", EmitDefaultValue = false)]
        public DateTime StateChangedAt { get; set; }

        /// <summary>
        /// Traits represent an identity&#39;s traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in &#x60;schema_url&#x60;.
        /// </summary>
        /// <value>Traits represent an identity&#39;s traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in &#x60;schema_url&#x60;.</value>
        [DataMember(Name = "traits", IsRequired = true, EmitDefaultValue = true)]
        public Object Traits { get; set; }

        /// <summary>
        /// UpdatedAt is a helper struct field for gobuffalo.pop.
        /// </summary>
        /// <value>UpdatedAt is a helper struct field for gobuffalo.pop.</value>
        [DataMember(Name = "updated_at", EmitDefaultValue = false)]
        public DateTime UpdatedAt { get; set; }

        /// <summary>
        /// VerifiableAddresses contains all the addresses that can be verified by the user.
        /// </summary>
        /// <value>VerifiableAddresses contains all the addresses that can be verified by the user.</value>
        [DataMember(Name = "verifiable_addresses", EmitDefaultValue = false)]
        public List<ClientVerifiableIdentityAddress> VerifiableAddresses { get; set; }

        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class ClientIdentity {\n");
            sb.Append("  CreatedAt: ").Append(CreatedAt).Append("\n");
            sb.Append("  Credentials: ").Append(Credentials).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  MetadataAdmin: ").Append(MetadataAdmin).Append("\n");
            sb.Append("  MetadataPublic: ").Append(MetadataPublic).Append("\n");
            sb.Append("  RecoveryAddresses: ").Append(RecoveryAddresses).Append("\n");
            sb.Append("  SchemaId: ").Append(SchemaId).Append("\n");
            sb.Append("  SchemaUrl: ").Append(SchemaUrl).Append("\n");
            sb.Append("  State: ").Append(State).Append("\n");
            sb.Append("  StateChangedAt: ").Append(StateChangedAt).Append("\n");
            sb.Append("  Traits: ").Append(Traits).Append("\n");
            sb.Append("  UpdatedAt: ").Append(UpdatedAt).Append("\n");
            sb.Append("  VerifiableAddresses: ").Append(VerifiableAddresses).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as ClientIdentity);
        }

        /// <summary>
        /// Returns true if ClientIdentity instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientIdentity to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientIdentity input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.CreatedAt == input.CreatedAt ||
                    (this.CreatedAt != null &&
                    this.CreatedAt.Equals(input.CreatedAt))
                ) && 
                (
                    this.Credentials == input.Credentials ||
                    this.Credentials != null &&
                    input.Credentials != null &&
                    this.Credentials.SequenceEqual(input.Credentials)
                ) && 
                (
                    this.Id == input.Id ||
                    (this.Id != null &&
                    this.Id.Equals(input.Id))
                ) && 
                (
                    this.MetadataAdmin == input.MetadataAdmin ||
                    (this.MetadataAdmin != null &&
                    this.MetadataAdmin.Equals(input.MetadataAdmin))
                ) && 
                (
                    this.MetadataPublic == input.MetadataPublic ||
                    (this.MetadataPublic != null &&
                    this.MetadataPublic.Equals(input.MetadataPublic))
                ) && 
                (
                    this.RecoveryAddresses == input.RecoveryAddresses ||
                    this.RecoveryAddresses != null &&
                    input.RecoveryAddresses != null &&
                    this.RecoveryAddresses.SequenceEqual(input.RecoveryAddresses)
                ) && 
                (
                    this.SchemaId == input.SchemaId ||
                    (this.SchemaId != null &&
                    this.SchemaId.Equals(input.SchemaId))
                ) && 
                (
                    this.SchemaUrl == input.SchemaUrl ||
                    (this.SchemaUrl != null &&
                    this.SchemaUrl.Equals(input.SchemaUrl))
                ) && 
                (
                    this.State == input.State ||
                    (this.State != null &&
                    this.State.Equals(input.State))
                ) && 
                (
                    this.StateChangedAt == input.StateChangedAt ||
                    (this.StateChangedAt != null &&
                    this.StateChangedAt.Equals(input.StateChangedAt))
                ) && 
                (
                    this.Traits == input.Traits ||
                    (this.Traits != null &&
                    this.Traits.Equals(input.Traits))
                ) && 
                (
                    this.UpdatedAt == input.UpdatedAt ||
                    (this.UpdatedAt != null &&
                    this.UpdatedAt.Equals(input.UpdatedAt))
                ) && 
                (
                    this.VerifiableAddresses == input.VerifiableAddresses ||
                    this.VerifiableAddresses != null &&
                    input.VerifiableAddresses != null &&
                    this.VerifiableAddresses.SequenceEqual(input.VerifiableAddresses)
                )
                && (this.AdditionalProperties.Count == input.AdditionalProperties.Count && !this.AdditionalProperties.Except(input.AdditionalProperties).Any());
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.CreatedAt != null)
                {
                    hashCode = (hashCode * 59) + this.CreatedAt.GetHashCode();
                }
                if (this.Credentials != null)
                {
                    hashCode = (hashCode * 59) + this.Credentials.GetHashCode();
                }
                if (this.Id != null)
                {
                    hashCode = (hashCode * 59) + this.Id.GetHashCode();
                }
                if (this.MetadataAdmin != null)
                {
                    hashCode = (hashCode * 59) + this.MetadataAdmin.GetHashCode();
                }
                if (this.MetadataPublic != null)
                {
                    hashCode = (hashCode * 59) + this.MetadataPublic.GetHashCode();
                }
                if (this.RecoveryAddresses != null)
                {
                    hashCode = (hashCode * 59) + this.RecoveryAddresses.GetHashCode();
                }
                if (this.SchemaId != null)
                {
                    hashCode = (hashCode * 59) + this.SchemaId.GetHashCode();
                }
                if (this.SchemaUrl != null)
                {
                    hashCode = (hashCode * 59) + this.SchemaUrl.GetHashCode();
                }
                if (this.State != null)
                {
                    hashCode = (hashCode * 59) + this.State.GetHashCode();
                }
                if (this.StateChangedAt != null)
                {
                    hashCode = (hashCode * 59) + this.StateChangedAt.GetHashCode();
                }
                if (this.Traits != null)
                {
                    hashCode = (hashCode * 59) + this.Traits.GetHashCode();
                }
                if (this.UpdatedAt != null)
                {
                    hashCode = (hashCode * 59) + this.UpdatedAt.GetHashCode();
                }
                if (this.VerifiableAddresses != null)
                {
                    hashCode = (hashCode * 59) + this.VerifiableAddresses.GetHashCode();
                }
                if (this.AdditionalProperties != null)
                {
                    hashCode = (hashCode * 59) + this.AdditionalProperties.GetHashCode();
                }
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
