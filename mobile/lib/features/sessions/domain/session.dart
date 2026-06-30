class Session {
  final String sessionId;
  final DateTime expiresAt;

  const Session({
    required this.sessionId,
    required this.expiresAt,
  });

  factory Session.fromJson(Map<String, dynamic> json) {
    return Session(
      sessionId: json["sessionId"] as String,
      expiresAt: DateTime.parse(json["expiresAt"] as String),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      "sessionId": sessionId,
      "expiresAt": expiresAt.toIso8601String(),
    };
  }
}