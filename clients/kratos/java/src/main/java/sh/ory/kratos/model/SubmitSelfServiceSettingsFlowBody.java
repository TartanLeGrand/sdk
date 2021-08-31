/*
 * Ory Kratos API
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.3-alpha.2
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.kratos.model;

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
import sh.ory.kratos.model.SubmitSelfServiceSettingsFlowWithOidcMethodBody;
import sh.ory.kratos.model.SubmitSelfServiceSettingsFlowWithPasswordMethodBody;
import sh.ory.kratos.model.SubmitSelfServiceSettingsFlowWithProfileMethodBody;

/**
 * SubmitSelfServiceSettingsFlowBody
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-08-31T13:45:14.536241677Z[Etc/UTC]")
public class SubmitSelfServiceSettingsFlowBody {
  public static final String SERIALIZED_NAME_CSRF_TOKEN = "csrf_token";
  @SerializedName(SERIALIZED_NAME_CSRF_TOKEN)
  private String csrfToken;

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  protected String method;

  public static final String SERIALIZED_NAME_PASSWORD = "password";
  @SerializedName(SERIALIZED_NAME_PASSWORD)
  private String password;

  public static final String SERIALIZED_NAME_TRAITS = "traits";
  @SerializedName(SERIALIZED_NAME_TRAITS)
  private Object traits;

  public SubmitSelfServiceSettingsFlowBody() {
    this.method = this.getClass().getSimpleName();
  }

  public SubmitSelfServiceSettingsFlowBody csrfToken(String csrfToken) {
    
    this.csrfToken = csrfToken;
    return this;
  }

   /**
   * The Anti-CSRF Token  This token is only required when performing browser flows.
   * @return csrfToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The Anti-CSRF Token  This token is only required when performing browser flows.")

  public String getCsrfToken() {
    return csrfToken;
  }


  public void setCsrfToken(String csrfToken) {
    this.csrfToken = csrfToken;
  }


  public SubmitSelfServiceSettingsFlowBody method(String method) {
    
    this.method = method;
    return this;
  }

   /**
   * Method  Should be set to profile when trying to update a profile.
   * @return method
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Method  Should be set to profile when trying to update a profile.")

  public String getMethod() {
    return method;
  }


  public void setMethod(String method) {
    this.method = method;
  }


  public SubmitSelfServiceSettingsFlowBody password(String password) {
    
    this.password = password;
    return this;
  }

   /**
   * Password is the updated password
   * @return password
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Password is the updated password")

  public String getPassword() {
    return password;
  }


  public void setPassword(String password) {
    this.password = password;
  }


  public SubmitSelfServiceSettingsFlowBody traits(Object traits) {
    
    this.traits = traits;
    return this;
  }

   /**
   * Traits contains all of the identity&#39;s traits.
   * @return traits
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Traits contains all of the identity's traits.")

  public Object getTraits() {
    return traits;
  }


  public void setTraits(Object traits) {
    this.traits = traits;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SubmitSelfServiceSettingsFlowBody submitSelfServiceSettingsFlowBody = (SubmitSelfServiceSettingsFlowBody) o;
    return Objects.equals(this.csrfToken, submitSelfServiceSettingsFlowBody.csrfToken) &&
        Objects.equals(this.method, submitSelfServiceSettingsFlowBody.method) &&
        Objects.equals(this.password, submitSelfServiceSettingsFlowBody.password) &&
        Objects.equals(this.traits, submitSelfServiceSettingsFlowBody.traits);
  }

  @Override
  public int hashCode() {
    return Objects.hash(csrfToken, method, password, traits);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SubmitSelfServiceSettingsFlowBody {\n");
    sb.append("    csrfToken: ").append(toIndentedString(csrfToken)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    password: ").append(toIndentedString(password)).append("\n");
    sb.append("    traits: ").append(toIndentedString(traits)).append("\n");
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

}

