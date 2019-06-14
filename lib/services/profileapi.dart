import 'dart:convert';
import 'dart:html';

import '../pathlookup.dart';

Future<String> createProfile(Object data) async {
  var url = await buildPath("Folio.API", "profile", new List<String>());
  return HttpRequest.requestCrossOrigin(url,
      method: "POST", sendData: jsonEncode(data));
}

void updateProfile(Object data, Function callback) async {
  var url = await buildPath("Folio.API", "profile", new List<String>());

  final request = HttpRequest();
  request.open("PUT", url);
  request.setRequestHeader("Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd.listen((e) => requestComplete(request, callback));
  request.send(jsonEncode(data));
}
