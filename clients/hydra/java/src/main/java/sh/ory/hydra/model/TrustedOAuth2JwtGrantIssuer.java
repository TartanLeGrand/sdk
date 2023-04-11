/*
 * Ory Hydra API
 * Documentation for all of Ory Hydra's APIs. 
 *
 * The version of the OpenAPI document: v2.1.1
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.hydra.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;
import sh.ory.hydra.model.TrustedOAuth2JwtGrantJsonWebKey;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import sh.ory.hydra.JSON;

/**
 * OAuth2 JWT Bearer Grant Type Issuer Trust Relationship
 */
@ApiModel(description = "OAuth2 JWT Bearer Grant Type Issuer Trust Relationship")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-04-11T10:57:16.155985320Z[Etc/UTC]")
public class TrustedOAuth2JwtGrantIssuer {
  public static final String SERIALIZED_NAME_ALLOW_ANY_SUBJECT = "allow_any_subject";
  @SerializedName(SERIALIZED_NAME_ALLOW_ANY_SUBJECT)
  private Boolean allowAnySubject;

  public static final String SERIALIZED_NAME_CREATED_AT = "created_at";
  @SerializedName(SERIALIZED_NAME_CREATED_AT)
  private OffsetDateTime createdAt;

  public static final String SERIALIZED_NAME_EXPIRES_AT = "expires_at";
  @SerializedName(SERIALIZED_NAME_EXPIRES_AT)
  private OffsetDateTime expiresAt;

  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private String id;

  public static final String SERIALIZED_NAME_ISSUER = "issuer";
  @SerializedName(SERIALIZED_NAME_ISSUER)
  private String issuer;

  public static final String SERIALIZED_NAME_PUBLIC_KEY = "public_key";
  @SerializedName(SERIALIZED_NAME_PUBLIC_KEY)
  private TrustedOAuth2JwtGrantJsonWebKey publicKey;

  public static final String SERIALIZED_NAME_SCOPE = "scope";
  @SerializedName(SERIALIZED_NAME_SCOPE)
  private List<String> scope = null;

  public static final String SERIALIZED_NAME_SUBJECT = "subject";
  @SerializedName(SERIALIZED_NAME_SUBJECT)
  private String subject;

  public TrustedOAuth2JwtGrantIssuer() {
  }

  public TrustedOAuth2JwtGrantIssuer allowAnySubject(Boolean allowAnySubject) {
    
    this.allowAnySubject = allowAnySubject;
    return this;
  }

   /**
   * The \&quot;allow_any_subject\&quot; indicates that the issuer is allowed to have any principal as the subject of the JWT.
   * @return allowAnySubject
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The \"allow_any_subject\" indicates that the issuer is allowed to have any principal as the subject of the JWT.")

  public Boolean getAllowAnySubject() {
    return allowAnySubject;
  }


  public void setAllowAnySubject(Boolean allowAnySubject) {
    this.allowAnySubject = allowAnySubject;
  }


  public TrustedOAuth2JwtGrantIssuer createdAt(OffsetDateTime createdAt) {
    
    this.createdAt = createdAt;
    return this;
  }

   /**
   * The \&quot;created_at\&quot; indicates, when grant was created.
   * @return createdAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The \"created_at\" indicates, when grant was created.")

  public OffsetDateTime getCreatedAt() {
    return createdAt;
  }


  public void setCreatedAt(OffsetDateTime createdAt) {
    this.createdAt = createdAt;
  }


  public TrustedOAuth2JwtGrantIssuer expiresAt(OffsetDateTime expiresAt) {
    
    this.expiresAt = expiresAt;
    return this;
  }

   /**
   * The \&quot;expires_at\&quot; indicates, when grant will expire, so we will reject assertion from \&quot;issuer\&quot; targeting \&quot;subject\&quot;.
   * @return expiresAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The \"expires_at\" indicates, when grant will expire, so we will reject assertion from \"issuer\" targeting \"subject\".")

  public OffsetDateTime getExpiresAt() {
    return expiresAt;
  }


  public void setExpiresAt(OffsetDateTime expiresAt) {
    this.expiresAt = expiresAt;
  }


  public TrustedOAuth2JwtGrantIssuer id(String id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "9edc811f-4e28-453c-9b46-4de65f00217f", value = "")

  public String getId() {
    return id;
  }


  public void setId(String id) {
    this.id = id;
  }


  public TrustedOAuth2JwtGrantIssuer issuer(String issuer) {
    
    this.issuer = issuer;
    return this;
  }

   /**
   * The \&quot;issuer\&quot; identifies the principal that issued the JWT assertion (same as \&quot;iss\&quot; claim in JWT).
   * @return issuer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "https://jwt-idp.example.com", value = "The \"issuer\" identifies the principal that issued the JWT assertion (same as \"iss\" claim in JWT).")

  public String getIssuer() {
    return issuer;
  }


  public void setIssuer(String issuer) {
    this.issuer = issuer;
  }


  public TrustedOAuth2JwtGrantIssuer publicKey(TrustedOAuth2JwtGrantJsonWebKey publicKey) {
    
    this.publicKey = publicKey;
    return this;
  }

   /**
   * Get publicKey
   * @return publicKey
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public TrustedOAuth2JwtGrantJsonWebKey getPublicKey() {
    return publicKey;
  }


  public void setPublicKey(TrustedOAuth2JwtGrantJsonWebKey publicKey) {
    this.publicKey = publicKey;
  }


  public TrustedOAuth2JwtGrantIssuer scope(List<String> scope) {
    
    this.scope = scope;
    return this;
  }

  public TrustedOAuth2JwtGrantIssuer addScopeItem(String scopeItem) {
    if (this.scope == null) {
      this.scope = new ArrayList<>();
    }
    this.scope.add(scopeItem);
    return this;
  }

   /**
   * The \&quot;scope\&quot; contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])
   * @return scope
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "[\"openid\",\"offline\"]", value = "The \"scope\" contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])")

  public List<String> getScope() {
    return scope;
  }


  public void setScope(List<String> scope) {
    this.scope = scope;
  }


  public TrustedOAuth2JwtGrantIssuer subject(String subject) {
    
    this.subject = subject;
    return this;
  }

   /**
   * The \&quot;subject\&quot; identifies the principal that is the subject of the JWT.
   * @return subject
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "mike@example.com", value = "The \"subject\" identifies the principal that is the subject of the JWT.")

  public String getSubject() {
    return subject;
  }


  public void setSubject(String subject) {
    this.subject = subject;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TrustedOAuth2JwtGrantIssuer trustedOAuth2JwtGrantIssuer = (TrustedOAuth2JwtGrantIssuer) o;
    return Objects.equals(this.allowAnySubject, trustedOAuth2JwtGrantIssuer.allowAnySubject) &&
        Objects.equals(this.createdAt, trustedOAuth2JwtGrantIssuer.createdAt) &&
        Objects.equals(this.expiresAt, trustedOAuth2JwtGrantIssuer.expiresAt) &&
        Objects.equals(this.id, trustedOAuth2JwtGrantIssuer.id) &&
        Objects.equals(this.issuer, trustedOAuth2JwtGrantIssuer.issuer) &&
        Objects.equals(this.publicKey, trustedOAuth2JwtGrantIssuer.publicKey) &&
        Objects.equals(this.scope, trustedOAuth2JwtGrantIssuer.scope) &&
        Objects.equals(this.subject, trustedOAuth2JwtGrantIssuer.subject);
  }

  @Override
  public int hashCode() {
    return Objects.hash(allowAnySubject, createdAt, expiresAt, id, issuer, publicKey, scope, subject);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TrustedOAuth2JwtGrantIssuer {\n");
    sb.append("    allowAnySubject: ").append(toIndentedString(allowAnySubject)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    expiresAt: ").append(toIndentedString(expiresAt)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    issuer: ").append(toIndentedString(issuer)).append("\n");
    sb.append("    publicKey: ").append(toIndentedString(publicKey)).append("\n");
    sb.append("    scope: ").append(toIndentedString(scope)).append("\n");
    sb.append("    subject: ").append(toIndentedString(subject)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("allow_any_subject");
    openapiFields.add("created_at");
    openapiFields.add("expires_at");
    openapiFields.add("id");
    openapiFields.add("issuer");
    openapiFields.add("public_key");
    openapiFields.add("scope");
    openapiFields.add("subject");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to TrustedOAuth2JwtGrantIssuer
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!TrustedOAuth2JwtGrantIssuer.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in TrustedOAuth2JwtGrantIssuer is not found in the empty JSON string", TrustedOAuth2JwtGrantIssuer.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!TrustedOAuth2JwtGrantIssuer.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `TrustedOAuth2JwtGrantIssuer` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("id") != null && !jsonObj.get("id").isJsonNull()) && !jsonObj.get("id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("id").toString()));
      }
      if ((jsonObj.get("issuer") != null && !jsonObj.get("issuer").isJsonNull()) && !jsonObj.get("issuer").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `issuer` to be a primitive type in the JSON string but got `%s`", jsonObj.get("issuer").toString()));
      }
      // validate the optional field `public_key`
      if (jsonObj.get("public_key") != null && !jsonObj.get("public_key").isJsonNull()) {
        TrustedOAuth2JwtGrantJsonWebKey.validateJsonObject(jsonObj.getAsJsonObject("public_key"));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("scope") != null && !jsonObj.get("scope").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `scope` to be an array in the JSON string but got `%s`", jsonObj.get("scope").toString()));
      }
      if ((jsonObj.get("subject") != null && !jsonObj.get("subject").isJsonNull()) && !jsonObj.get("subject").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `subject` to be a primitive type in the JSON string but got `%s`", jsonObj.get("subject").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!TrustedOAuth2JwtGrantIssuer.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'TrustedOAuth2JwtGrantIssuer' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<TrustedOAuth2JwtGrantIssuer> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(TrustedOAuth2JwtGrantIssuer.class));

       return (TypeAdapter<T>) new TypeAdapter<TrustedOAuth2JwtGrantIssuer>() {
           @Override
           public void write(JsonWriter out, TrustedOAuth2JwtGrantIssuer value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public TrustedOAuth2JwtGrantIssuer read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of TrustedOAuth2JwtGrantIssuer given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of TrustedOAuth2JwtGrantIssuer
  * @throws IOException if the JSON string is invalid with respect to TrustedOAuth2JwtGrantIssuer
  */
  public static TrustedOAuth2JwtGrantIssuer fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, TrustedOAuth2JwtGrantIssuer.class);
  }

 /**
  * Convert an instance of TrustedOAuth2JwtGrantIssuer to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

