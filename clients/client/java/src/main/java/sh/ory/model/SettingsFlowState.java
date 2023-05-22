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
import io.swagger.annotations.ApiModel;
import com.google.gson.annotations.SerializedName;

import java.io.IOException;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;

/**
 * show_form: No user data has been collected, or it is invalid, and thus the form should be shown. success: Indicates that the settings flow has been updated successfully with the provided data. Done will stay true when repeatedly checking. If set to true, done will revert back to false only when a flow with invalid (e.g. \&quot;please use a valid phone number\&quot;) data was sent.
 */
@JsonAdapter(SettingsFlowState.Adapter.class)
public enum SettingsFlowState {
  
  SHOW_FORM("show_form"),
  
  SUCCESS("success");

  private String value;

  SettingsFlowState(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static SettingsFlowState fromValue(String value) {
    for (SettingsFlowState b : SettingsFlowState.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<SettingsFlowState> {
    @Override
    public void write(final JsonWriter jsonWriter, final SettingsFlowState enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public SettingsFlowState read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return SettingsFlowState.fromValue(value);
    }
  }
}

