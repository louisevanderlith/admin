import 'dart:convert';
import 'dart:html';

import 'pathlookup.dart';

Future updateProfile(Object data) async {
  var url = await buildPath("Folio.API", "message", new List<String>());
  var resp = await HttpRequest.requestCrossOrigin(url,
      method: "POST", sendData: jsonEncode(data));

  print(resp);
}
