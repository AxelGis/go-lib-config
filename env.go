package config

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// FillStructFromEnv recursively fills a struct with values from environment variables based on struct tags.
func FillStructFromEnv(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		// If the field is a nested struct, recursively fill it
		if field.Kind() == reflect.Struct {
			FillStructFromEnv(field)
			continue
		}

		envVar, hasEnvTag := fieldType.Tag.Lookup("env")
		if !hasEnvTag {
			continue
		}

		envValue := getEnvWithDefault(envVar, fieldType.Tag.Get("default"))
		setFieldValueFromEnv(field, envValue)
	}
}

// getEnvWithDefault tries to get an environment variable; returns default value if not found.
func getEnvWithDefault(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// setFieldValueFromEnv sets the field value based on its kind and the environment variable value.
func setFieldValueFromEnv(field reflect.Value, envValue string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(envValue)
	case reflect.Int, reflect.Int64:
		if field.Type() == reflect.TypeOf(time.Duration(0)) {
			if duration, err := time.ParseDuration(envValue); err == nil {
				field.Set(reflect.ValueOf(duration))
			}
		} else if intValue, err := strconv.ParseInt(envValue, 10, 64); err == nil {
			field.SetInt(intValue)
		}
	case reflect.Uint32:
		if intValue, err := strconv.ParseUint(envValue, 10, 32); err == nil {
			field.SetUint(intValue)
		}
	case reflect.Bool:
		if boolValue, err := strconv.ParseBool(envValue); err == nil {
			field.SetBool(boolValue)
		}
	case reflect.Slice:
		setSliceFromEnv(field, envValue)
	case reflect.Float64, reflect.Float32:
		if floatValue, err := strconv.ParseFloat(envValue, 64); err == nil {
			field.SetFloat(floatValue)
		}
	default:
		if field.Type() == reflect.TypeOf(time.Duration(0)) {
			if duration, err := time.ParseDuration(envValue); err == nil {
				field.Set(reflect.ValueOf(duration))
			}
		}
	}
}

// setSliceFromEnv sets a slice field from a comma-separated environment variable.
func setSliceFromEnv(field reflect.Value, envValue string) {
	elemKind := field.Type().Elem().Kind()
	values := strings.Split(envValue, ",")
	slice := reflect.MakeSlice(field.Type(), len(values), len(values))
	for i, val := range values {
		val = strings.TrimSpace(val)
		switch elemKind {
		case reflect.String:
			slice.Index(i).SetString(val)
		case reflect.Int, reflect.Int64:
			if intValue, err := strconv.ParseInt(val, 10, 64); err == nil {
				slice.Index(i).SetInt(intValue)
			}
		case reflect.Uint64:
			if uintValue, err := strconv.ParseUint(val, 10, 64); err == nil {
				slice.Index(i).SetUint(uintValue)
			}
		case reflect.Bool:
			if boolValue, err := strconv.ParseBool(val); err == nil {
				slice.Index(i).SetBool(boolValue)
			}
		}
	}
	field.Set(slice)
}

// GetEnv returns the environment variable value or a default value if not set.
func GetEnv(key, defaultVal string) string {
	return getEnvWithDefault(key, defaultVal)
}

// GetEnvAsInt returns the environment variable as an integer or a default value if not set or conversion fails.
func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

// GetEnvAsBool returns the environment variable as a boolean or a default value if not set or conversion fails.
func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := GetEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

// GetEnvAsSlice returns the environment variable as a string slice or a default value if not set.
func GetEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := GetEnv(name, "")
	if valStr == "" {
		return defaultVal
	}
	return strings.Split(valStr, sep)
}
