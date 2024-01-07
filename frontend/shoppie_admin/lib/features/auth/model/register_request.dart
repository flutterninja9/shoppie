import 'package:json_annotation/json_annotation.dart';

part "register_request.g.dart";

@JsonSerializable()
class RegisterRequest {
  final String username;
  @JsonKey(name: "first_name")
  final String firstName;
  @JsonKey(name: "last_name")
  final String lastName;
  final String email;
  final String password;

  RegisterRequest({
    required this.username,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.password,
  });

  factory RegisterRequest.fromJson(Map<String, dynamic> json) =>
      _$RegisterRequestFromJson(json);

  Map<String, dynamic> toJson() => _$RegisterRequestToJson(this);
}
