/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.14
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
    /// ClientExpandedPermissionTree
    /// </summary>
    [DataContract(Name = "expandedPermissionTree")]
    public partial class ClientExpandedPermissionTree : IEquatable<ClientExpandedPermissionTree>, IValidatableObject
    {
        /// <summary>
        /// The type of the node. union TreeNodeUnion exclusion TreeNodeExclusion intersection TreeNodeIntersection leaf TreeNodeLeaf tuple_to_subject_set TreeNodeTupleToSubjectSet computed_subject_set TreeNodeComputedSubjectSet not TreeNodeNot unspecified TreeNodeUnspecified
        /// </summary>
        /// <value>The type of the node. union TreeNodeUnion exclusion TreeNodeExclusion intersection TreeNodeIntersection leaf TreeNodeLeaf tuple_to_subject_set TreeNodeTupleToSubjectSet computed_subject_set TreeNodeComputedSubjectSet not TreeNodeNot unspecified TreeNodeUnspecified</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum TypeEnum
        {
            /// <summary>
            /// Enum Union for value: union
            /// </summary>
            [EnumMember(Value = "union")]
            Union = 1,

            /// <summary>
            /// Enum Exclusion for value: exclusion
            /// </summary>
            [EnumMember(Value = "exclusion")]
            Exclusion = 2,

            /// <summary>
            /// Enum Intersection for value: intersection
            /// </summary>
            [EnumMember(Value = "intersection")]
            Intersection = 3,

            /// <summary>
            /// Enum Leaf for value: leaf
            /// </summary>
            [EnumMember(Value = "leaf")]
            Leaf = 4,

            /// <summary>
            /// Enum TupleToSubjectSet for value: tuple_to_subject_set
            /// </summary>
            [EnumMember(Value = "tuple_to_subject_set")]
            TupleToSubjectSet = 5,

            /// <summary>
            /// Enum ComputedSubjectSet for value: computed_subject_set
            /// </summary>
            [EnumMember(Value = "computed_subject_set")]
            ComputedSubjectSet = 6,

            /// <summary>
            /// Enum Not for value: not
            /// </summary>
            [EnumMember(Value = "not")]
            Not = 7,

            /// <summary>
            /// Enum Unspecified for value: unspecified
            /// </summary>
            [EnumMember(Value = "unspecified")]
            Unspecified = 8

        }


        /// <summary>
        /// The type of the node. union TreeNodeUnion exclusion TreeNodeExclusion intersection TreeNodeIntersection leaf TreeNodeLeaf tuple_to_subject_set TreeNodeTupleToSubjectSet computed_subject_set TreeNodeComputedSubjectSet not TreeNodeNot unspecified TreeNodeUnspecified
        /// </summary>
        /// <value>The type of the node. union TreeNodeUnion exclusion TreeNodeExclusion intersection TreeNodeIntersection leaf TreeNodeLeaf tuple_to_subject_set TreeNodeTupleToSubjectSet computed_subject_set TreeNodeComputedSubjectSet not TreeNodeNot unspecified TreeNodeUnspecified</value>
        [DataMember(Name = "type", IsRequired = true, EmitDefaultValue = false)]
        public TypeEnum Type { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientExpandedPermissionTree" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientExpandedPermissionTree()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientExpandedPermissionTree" /> class.
        /// </summary>
        /// <param name="children">The children of the node, possibly none..</param>
        /// <param name="tuple">tuple.</param>
        /// <param name="type">The type of the node. union TreeNodeUnion exclusion TreeNodeExclusion intersection TreeNodeIntersection leaf TreeNodeLeaf tuple_to_subject_set TreeNodeTupleToSubjectSet computed_subject_set TreeNodeComputedSubjectSet not TreeNodeNot unspecified TreeNodeUnspecified (required).</param>
        public ClientExpandedPermissionTree(List<ClientExpandedPermissionTree> children = default(List<ClientExpandedPermissionTree>), ClientRelationship tuple = default(ClientRelationship), TypeEnum type = default(TypeEnum))
        {
            this.Type = type;
            this.Children = children;
            this.Tuple = tuple;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The children of the node, possibly none.
        /// </summary>
        /// <value>The children of the node, possibly none.</value>
        [DataMember(Name = "children", EmitDefaultValue = false)]
        public List<ClientExpandedPermissionTree> Children { get; set; }

        /// <summary>
        /// Gets or Sets Tuple
        /// </summary>
        [DataMember(Name = "tuple", EmitDefaultValue = false)]
        public ClientRelationship Tuple { get; set; }

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
            sb.Append("class ClientExpandedPermissionTree {\n");
            sb.Append("  Children: ").Append(Children).Append("\n");
            sb.Append("  Tuple: ").Append(Tuple).Append("\n");
            sb.Append("  Type: ").Append(Type).Append("\n");
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
            return this.Equals(input as ClientExpandedPermissionTree);
        }

        /// <summary>
        /// Returns true if ClientExpandedPermissionTree instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientExpandedPermissionTree to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientExpandedPermissionTree input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Children == input.Children ||
                    this.Children != null &&
                    input.Children != null &&
                    this.Children.SequenceEqual(input.Children)
                ) && 
                (
                    this.Tuple == input.Tuple ||
                    (this.Tuple != null &&
                    this.Tuple.Equals(input.Tuple))
                ) && 
                (
                    this.Type == input.Type ||
                    this.Type.Equals(input.Type)
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
                if (this.Children != null)
                {
                    hashCode = (hashCode * 59) + this.Children.GetHashCode();
                }
                if (this.Tuple != null)
                {
                    hashCode = (hashCode * 59) + this.Tuple.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Type.GetHashCode();
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
