/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.18
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
    /// ClientInternalRelationTuple
    /// </summary>
    [DataContract(Name = "InternalRelationTuple")]
    public partial class ClientInternalRelationTuple : IEquatable<ClientInternalRelationTuple>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientInternalRelationTuple" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientInternalRelationTuple()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientInternalRelationTuple" /> class.
        /// </summary>
        /// <param name="_namespace">Namespace of the Relation Tuple (required).</param>
        /// <param name="_object">Object of the Relation Tuple (required).</param>
        /// <param name="relation">Relation of the Relation Tuple (required).</param>
        /// <param name="subjectId">SubjectID of the Relation Tuple  Either SubjectSet or SubjectID are required..</param>
        /// <param name="subjectSet">subjectSet.</param>
        public ClientInternalRelationTuple(string _namespace = default(string), string _object = default(string), string relation = default(string), string subjectId = default(string), ClientSubjectSet subjectSet = default(ClientSubjectSet))
        {
            // to ensure "_namespace" is required (not null)
            if (_namespace == null) {
                throw new ArgumentNullException("_namespace is a required property for ClientInternalRelationTuple and cannot be null");
            }
            this.Namespace = _namespace;
            // to ensure "_object" is required (not null)
            if (_object == null) {
                throw new ArgumentNullException("_object is a required property for ClientInternalRelationTuple and cannot be null");
            }
            this.Object = _object;
            // to ensure "relation" is required (not null)
            if (relation == null) {
                throw new ArgumentNullException("relation is a required property for ClientInternalRelationTuple and cannot be null");
            }
            this.Relation = relation;
            this.SubjectId = subjectId;
            this.SubjectSet = subjectSet;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Namespace of the Relation Tuple
        /// </summary>
        /// <value>Namespace of the Relation Tuple</value>
        [DataMember(Name = "namespace", IsRequired = true, EmitDefaultValue = false)]
        public string Namespace { get; set; }

        /// <summary>
        /// Object of the Relation Tuple
        /// </summary>
        /// <value>Object of the Relation Tuple</value>
        [DataMember(Name = "object", IsRequired = true, EmitDefaultValue = false)]
        public string Object { get; set; }

        /// <summary>
        /// Relation of the Relation Tuple
        /// </summary>
        /// <value>Relation of the Relation Tuple</value>
        [DataMember(Name = "relation", IsRequired = true, EmitDefaultValue = false)]
        public string Relation { get; set; }

        /// <summary>
        /// SubjectID of the Relation Tuple  Either SubjectSet or SubjectID are required.
        /// </summary>
        /// <value>SubjectID of the Relation Tuple  Either SubjectSet or SubjectID are required.</value>
        [DataMember(Name = "subject_id", EmitDefaultValue = false)]
        public string SubjectId { get; set; }

        /// <summary>
        /// Gets or Sets SubjectSet
        /// </summary>
        [DataMember(Name = "subject_set", EmitDefaultValue = false)]
        public ClientSubjectSet SubjectSet { get; set; }

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
            sb.Append("class ClientInternalRelationTuple {\n");
            sb.Append("  Namespace: ").Append(Namespace).Append("\n");
            sb.Append("  Object: ").Append(Object).Append("\n");
            sb.Append("  Relation: ").Append(Relation).Append("\n");
            sb.Append("  SubjectId: ").Append(SubjectId).Append("\n");
            sb.Append("  SubjectSet: ").Append(SubjectSet).Append("\n");
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
            return this.Equals(input as ClientInternalRelationTuple);
        }

        /// <summary>
        /// Returns true if ClientInternalRelationTuple instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientInternalRelationTuple to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientInternalRelationTuple input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Namespace == input.Namespace ||
                    (this.Namespace != null &&
                    this.Namespace.Equals(input.Namespace))
                ) && 
                (
                    this.Object == input.Object ||
                    (this.Object != null &&
                    this.Object.Equals(input.Object))
                ) && 
                (
                    this.Relation == input.Relation ||
                    (this.Relation != null &&
                    this.Relation.Equals(input.Relation))
                ) && 
                (
                    this.SubjectId == input.SubjectId ||
                    (this.SubjectId != null &&
                    this.SubjectId.Equals(input.SubjectId))
                ) && 
                (
                    this.SubjectSet == input.SubjectSet ||
                    (this.SubjectSet != null &&
                    this.SubjectSet.Equals(input.SubjectSet))
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
                if (this.Namespace != null)
                {
                    hashCode = (hashCode * 59) + this.Namespace.GetHashCode();
                }
                if (this.Object != null)
                {
                    hashCode = (hashCode * 59) + this.Object.GetHashCode();
                }
                if (this.Relation != null)
                {
                    hashCode = (hashCode * 59) + this.Relation.GetHashCode();
                }
                if (this.SubjectId != null)
                {
                    hashCode = (hashCode * 59) + this.SubjectId.GetHashCode();
                }
                if (this.SubjectSet != null)
                {
                    hashCode = (hashCode * 59) + this.SubjectSet.GetHashCode();
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
