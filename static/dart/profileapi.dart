import 'dart:convert';
import 'dart:html';

import 'pathlookup.dart';

Future<String> createProfile(Object data) async {
  var url = await buildPath("Folio.API", "profile", new List<String>());
  return HttpRequest.requestCrossOrigin(url,
      method: "POST", sendData: jsonEncode(data));
}

Future<String> updateProfile(Object data) async {
  var url = await buildPath("Folio.API", "profile", new List<String>());
  return HttpRequest.requestCrossOrigin(url,
      method: "PUT", sendData: jsonEncode(data));
}