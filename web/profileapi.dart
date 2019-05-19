import 'dart:convert';
import 'dart:html';

import 'pathlookup.dart';

Future<String> createProfile(Object data) async {
  var url = await buildPath("Folio.API", "profile", new List<String>());
  return HttpRequest.requestCrossOrigin(url,
      method: "POST", sendData: jsonEncode(data));
}

Future updateProfile(Object data, Function callback) async {
  var url = await buildPath("Folio.API", "profile", new List<String>());

  final request = HttpRequest();
  request.open("PUT", url);
  request.withCredentials = true;
  request.onLoadEnd.listen((e) => requestComplete(request, callback));
  request.send(jsonEncode(data));

  /* return HttpRequest.requestCrossOrigin(url,
      method: "PUT", sendData: jsonEncode(data));*/
}
