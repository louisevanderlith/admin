import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'pathlookup.dart';

Future<HttpRequest> createArticle(String title, String username) async {
  var url = await buildPath("Blog.API", "article", new List<String>());
  var data = jsonEncode({
    "Title": title,
    "ImageKey": '0`0',
    "Content": '',
    "WrittenBy": username,
    "Public": false
  });

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("POST", url);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

Future<HttpRequest> updateArticle(String title, String content, String imageKey, String username) async {
  var url = await buildPath("Folio.API", "profile", new List<String>());
  final data = jsonEncode({
    "Title": title,
    "ImageKey": imageKey,
    "Content": content,
    "WrittenBy": username,
    "Public": false
  });

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("PUT", url);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);
}

void onProgress(ProgressEvent e) {
  if (e.lengthComputable) {
    print('Progress... ${e.total}/${e.loaded}');
  }
}