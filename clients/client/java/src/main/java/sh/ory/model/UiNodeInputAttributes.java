/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.33
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.model;

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
import org.openapitools.jackson.nullable.JsonNullable;
import sh.ory.model.UiText;

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

import sh.ory.JSON;

/**
 * InputAttributes represents the attributes of an input node
 */
@ApiModel(description = "InputAttributes represents the attributes of an input node")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-05-22T13:18:59.927245844Z[Etc/UTC]")
public class UiNodeInputAttributes {
  /**
   * The autocomplete attribute for the input. email InputAttributeAutocompleteEmail tel InputAttributeAutocompleteTel url InputAttributeAutocompleteUrl current-password InputAttributeAutocompleteCurrentPassword new-password InputAttributeAutocompleteNewPassword one-time-code InputAttributeAutocompleteOneTimeCode
   */
  @JsonAdapter(AutocompleteEnum.Adapter.class)
  public enum AutocompleteEnum {
    EMAIL("email"),
    
    TEL("tel"),
    
    URL("url"),
    
    CURRENT_PASSWORD("current-password"),
    
    NEW_PASSWORD("new-password"),
    
    ONE_TIME_CODE("one-time-code");

    private String value;

    AutocompleteEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static AutocompleteEnum fromValue(String value) {
      for (AutocompleteEnum b : AutocompleteEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<AutocompleteEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final AutocompleteEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public AutocompleteEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return AutocompleteEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_AUTOCOMPLETE = "autocomplete";
  @SerializedName(SERIALIZED_NAME_AUTOCOMPLETE)
  private AutocompleteEnum autocomplete;

  public static final String SERIALIZED_NAME_DISABLED = "disabled";
  @SerializedName(SERIALIZED_NAME_DISABLED)
  private Boolean disabled;

  public static final String SERIALIZED_NAME_LABEL = "label";
  @SerializedName(SERIALIZED_NAME_LABEL)
  private UiText label;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_NODE_TYPE = "node_type";
  @SerializedName(SERIALIZED_NAME_NODE_TYPE)
  private String nodeType;

  public static final String SERIALIZED_NAME_ONCLICK = "onclick";
  @SerializedName(SERIALIZED_NAME_ONCLICK)
  private String onclick;

  public static final String SERIALIZED_NAME_PATTERN = "pattern";
  @SerializedName(SERIALIZED_NAME_PATTERN)
  private String pattern;

  public static final String SERIALIZED_NAME_REQUIRED = "required";
  @SerializedName(SERIALIZED_NAME_REQUIRED)
  private Boolean required;

  /**
   * The input&#39;s element type. text InputAttributeTypeText password InputAttributeTypePassword number InputAttributeTypeNumber checkbox InputAttributeTypeCheckbox hidden InputAttributeTypeHidden email InputAttributeTypeEmail tel InputAttributeTypeTel submit InputAttributeTypeSubmit button InputAttributeTypeButton datetime-local InputAttributeTypeDateTimeLocal date InputAttributeTypeDate url InputAttributeTypeURI
   */
  @JsonAdapter(TypeEnum.Adapter.class)
  public enum TypeEnum {
    TEXT("text"),
    
    PASSWORD("password"),
    
    NUMBER("number"),
    
    CHECKBOX("checkbox"),
    
    HIDDEN("hidden"),
    
    EMAIL("email"),
    
    TEL("tel"),
    
    SUBMIT("submit"),
    
    BUTTON("button"),
    
    DATETIME_LOCAL("datetime-local"),
    
    DATE("date"),
    
    URL("url");

    private String value;

    TypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static TypeEnum fromValue(String value) {
      for (TypeEnum b : TypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<TypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final TypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public TypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return TypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private TypeEnum type;

  public static final String SERIALIZED_NAME_VALUE = "value";
  @SerializedName(SERIALIZED_NAME_VALUE)
  private Object value = null;

  public UiNodeInputAttributes() {
  }

  public UiNodeInputAttributes autocomplete(AutocompleteEnum autocomplete) {
    
    this.autocomplete = autocomplete;
    return this;
  }

   /**
   * The autocomplete attribute for the input. email InputAttributeAutocompleteEmail tel InputAttributeAutocompleteTel url InputAttributeAutocompleteUrl current-password InputAttributeAutocompleteCurrentPassword new-password InputAttributeAutocompleteNewPassword one-time-code InputAttributeAutocompleteOneTimeCode
   * @return autocomplete
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The autocomplete attribute for the input. email InputAttributeAutocompleteEmail tel InputAttributeAutocompleteTel url InputAttributeAutocompleteUrl current-password InputAttributeAutocompleteCurrentPassword new-password InputAttributeAutocompleteNewPassword one-time-code InputAttributeAutocompleteOneTimeCode")

  public AutocompleteEnum getAutocomplete() {
    return autocomplete;
  }


  public void setAutocomplete(AutocompleteEnum autocomplete) {
    this.autocomplete = autocomplete;
  }


  public UiNodeInputAttributes disabled(Boolean disabled) {
    
    this.disabled = disabled;
    return this;
  }

   /**
   * Sets the input&#39;s disabled field to true or false.
   * @return disabled
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Sets the input's disabled field to true or false.")

  public Boolean getDisabled() {
    return disabled;
  }


  public void setDisabled(Boolean disabled) {
    this.disabled = disabled;
  }


  public UiNodeInputAttributes label(UiText label) {
    
    this.label = label;
    return this;
  }

   /**
   * Get label
   * @return label
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public UiText getLabel() {
    return label;
  }


  public void setLabel(UiText label) {
    this.label = label;
  }


  public UiNodeInputAttributes name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * The input&#39;s element name.
   * @return name
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The input's element name.")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public UiNodeInputAttributes nodeType(String nodeType) {
    
    this.nodeType = nodeType;
    return this;
  }

   /**
   * NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \&quot;input\&quot;.
   * @return nodeType
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "NodeType represents this node's types. It is a mirror of `node.type` and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \"input\".")

  public String getNodeType() {
    return nodeType;
  }


  public void setNodeType(String nodeType) {
    this.nodeType = nodeType;
  }


  public UiNodeInputAttributes onclick(String onclick) {
    
    this.onclick = onclick;
    return this;
  }

   /**
   * OnClick may contain javascript which should be executed on click. This is primarily used for WebAuthn.
   * @return onclick
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "OnClick may contain javascript which should be executed on click. This is primarily used for WebAuthn.")

  public String getOnclick() {
    return onclick;
  }


  public void setOnclick(String onclick) {
    this.onclick = onclick;
  }


  public UiNodeInputAttributes pattern(String pattern) {
    
    this.pattern = pattern;
    return this;
  }

   /**
   * The input&#39;s pattern.
   * @return pattern
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The input's pattern.")

  public String getPattern() {
    return pattern;
  }


  public void setPattern(String pattern) {
    this.pattern = pattern;
  }


  public UiNodeInputAttributes required(Boolean required) {
    
    this.required = required;
    return this;
  }

   /**
   * Mark this input field as required.
   * @return required
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Mark this input field as required.")

  public Boolean getRequired() {
    return required;
  }


  public void setRequired(Boolean required) {
    this.required = required;
  }


  public UiNodeInputAttributes type(TypeEnum type) {
    
    this.type = type;
    return this;
  }

   /**
   * The input&#39;s element type. text InputAttributeTypeText password InputAttributeTypePassword number InputAttributeTypeNumber checkbox InputAttributeTypeCheckbox hidden InputAttributeTypeHidden email InputAttributeTypeEmail tel InputAttributeTypeTel submit InputAttributeTypeSubmit button InputAttributeTypeButton datetime-local InputAttributeTypeDateTimeLocal date InputAttributeTypeDate url InputAttributeTypeURI
   * @return type
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The input's element type. text InputAttributeTypeText password InputAttributeTypePassword number InputAttributeTypeNumber checkbox InputAttributeTypeCheckbox hidden InputAttributeTypeHidden email InputAttributeTypeEmail tel InputAttributeTypeTel submit InputAttributeTypeSubmit button InputAttributeTypeButton datetime-local InputAttributeTypeDateTimeLocal date InputAttributeTypeDate url InputAttributeTypeURI")

  public TypeEnum getType() {
    return type;
  }


  public void setType(TypeEnum type) {
    this.type = type;
  }


  public UiNodeInputAttributes value(Object value) {
    
    this.value = value;
    return this;
  }

   /**
   * The input&#39;s value.
   * @return value
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The input's value.")

  public Object getValue() {
    return value;
  }


  public void setValue(Object value) {
    this.value = value;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UiNodeInputAttributes uiNodeInputAttributes = (UiNodeInputAttributes) o;
    return Objects.equals(this.autocomplete, uiNodeInputAttributes.autocomplete) &&
        Objects.equals(this.disabled, uiNodeInputAttributes.disabled) &&
        Objects.equals(this.label, uiNodeInputAttributes.label) &&
        Objects.equals(this.name, uiNodeInputAttributes.name) &&
        Objects.equals(this.nodeType, uiNodeInputAttributes.nodeType) &&
        Objects.equals(this.onclick, uiNodeInputAttributes.onclick) &&
        Objects.equals(this.pattern, uiNodeInputAttributes.pattern) &&
        Objects.equals(this.required, uiNodeInputAttributes.required) &&
        Objects.equals(this.type, uiNodeInputAttributes.type) &&
        Objects.equals(this.value, uiNodeInputAttributes.value);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(autocomplete, disabled, label, name, nodeType, onclick, pattern, required, type, value);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UiNodeInputAttributes {\n");
    sb.append("    autocomplete: ").append(toIndentedString(autocomplete)).append("\n");
    sb.append("    disabled: ").append(toIndentedString(disabled)).append("\n");
    sb.append("    label: ").append(toIndentedString(label)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    nodeType: ").append(toIndentedString(nodeType)).append("\n");
    sb.append("    onclick: ").append(toIndentedString(onclick)).append("\n");
    sb.append("    pattern: ").append(toIndentedString(pattern)).append("\n");
    sb.append("    required: ").append(toIndentedString(required)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
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
    openapiFields.add("autocomplete");
    openapiFields.add("disabled");
    openapiFields.add("label");
    openapiFields.add("name");
    openapiFields.add("node_type");
    openapiFields.add("onclick");
    openapiFields.add("pattern");
    openapiFields.add("required");
    openapiFields.add("type");
    openapiFields.add("value");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("disabled");
    openapiRequiredFields.add("name");
    openapiRequiredFields.add("node_type");
    openapiRequiredFields.add("type");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to UiNodeInputAttributes
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!UiNodeInputAttributes.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in UiNodeInputAttributes is not found in the empty JSON string", UiNodeInputAttributes.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!UiNodeInputAttributes.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `UiNodeInputAttributes` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : UiNodeInputAttributes.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if ((jsonObj.get("autocomplete") != null && !jsonObj.get("autocomplete").isJsonNull()) && !jsonObj.get("autocomplete").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `autocomplete` to be a primitive type in the JSON string but got `%s`", jsonObj.get("autocomplete").toString()));
      }
      // validate the optional field `label`
      if (jsonObj.get("label") != null && !jsonObj.get("label").isJsonNull()) {
        UiText.validateJsonObject(jsonObj.getAsJsonObject("label"));
      }
      if (!jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if (!jsonObj.get("node_type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `node_type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("node_type").toString()));
      }
      if ((jsonObj.get("onclick") != null && !jsonObj.get("onclick").isJsonNull()) && !jsonObj.get("onclick").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `onclick` to be a primitive type in the JSON string but got `%s`", jsonObj.get("onclick").toString()));
      }
      if ((jsonObj.get("pattern") != null && !jsonObj.get("pattern").isJsonNull()) && !jsonObj.get("pattern").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `pattern` to be a primitive type in the JSON string but got `%s`", jsonObj.get("pattern").toString()));
      }
      if (!jsonObj.get("type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("type").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!UiNodeInputAttributes.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'UiNodeInputAttributes' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<UiNodeInputAttributes> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(UiNodeInputAttributes.class));

       return (TypeAdapter<T>) new TypeAdapter<UiNodeInputAttributes>() {
           @Override
           public void write(JsonWriter out, UiNodeInputAttributes value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public UiNodeInputAttributes read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of UiNodeInputAttributes given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of UiNodeInputAttributes
  * @throws IOException if the JSON string is invalid with respect to UiNodeInputAttributes
  */
  public static UiNodeInputAttributes fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, UiNodeInputAttributes.class);
  }

 /**
  * Convert an instance of UiNodeInputAttributes to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

