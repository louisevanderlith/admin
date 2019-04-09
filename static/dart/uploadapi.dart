import 'dart:html';
import 'pathlookup.dart';

Future<HttpRequest> createUpload(FormData data) async {
  var path = await buildPath('Artifact.API', 'upload', new List<String>());

  return HttpRequest.request(path,
      method: 'POST',
      withCredentials: false,
      sendData: data);
}
