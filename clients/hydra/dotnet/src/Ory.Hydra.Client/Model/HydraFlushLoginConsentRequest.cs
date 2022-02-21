/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs. 
 *
 * The version of the OpenAPI document: v1.11.5
 * Contact: hi@ory.sh
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
using OpenAPIDateConverter = Ory.Hydra.Client.Client.OpenAPIDateConverter;

namespace Ory.Hydra.Client.Model
{
    /// <summary>
    /// HydraFlushLoginConsentRequest
    /// </summary>
    [DataContract(Name = "flushLoginConsentRequest")]
    public partial class HydraFlushLoginConsentRequest : IEquatable<HydraFlushLoginConsentRequest>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="HydraFlushLoginConsentRequest" /> class.
        /// </summary>
        /// <param name="notAfter">NotAfter sets after which point tokens should not be flushed. This is useful when you want to keep a history of recent login and consent database entries for auditing..</param>
        public HydraFlushLoginConsentRequest(DateTime notAfter = default(DateTime))
        {
            this.NotAfter = notAfter;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// NotAfter sets after which point tokens should not be flushed. This is useful when you want to keep a history of recent login and consent database entries for auditing.
        /// </summary>
        /// <value>NotAfter sets after which point tokens should not be flushed. This is useful when you want to keep a history of recent login and consent database entries for auditing.</value>
        [DataMember(Name = "notAfter", EmitDefaultValue = false)]
        public DateTime NotAfter { get; set; }

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
            sb.Append("class HydraFlushLoginConsentRequest {\n");
            sb.Append("  NotAfter: ").Append(NotAfter).Append("\n");
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
            return this.Equals(input as HydraFlushLoginConsentRequest);
        }

        /// <summary>
        /// Returns true if HydraFlushLoginConsentRequest instances are equal
        /// </summary>
        /// <param name="input">Instance of HydraFlushLoginConsentRequest to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(HydraFlushLoginConsentRequest input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.NotAfter == input.NotAfter ||
                    (this.NotAfter != null &&
                    this.NotAfter.Equals(input.NotAfter))
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
                if (this.NotAfter != null)
                {
                    hashCode = (hashCode * 59) + this.NotAfter.GetHashCode();
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