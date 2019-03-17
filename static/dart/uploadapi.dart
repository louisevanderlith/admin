import 'dart:convert';
import 'dart:html';

import 'pathlookup.dart';

void createUpload(data) {
  buildPath('Artifact.API', 'upload', new List<String>()).then((buildPath) => {
        HttpRequest.requestCrossOrigin(buildPath,
            method: "POST", sendData: jsonEncode(data))
      });
}
