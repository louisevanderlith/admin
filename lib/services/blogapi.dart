import 'dart:async';
import 'dart:convert';
import 'dart:html';

import '../pathlookup.dart';

Future<HttpRequest> createArticle(String title, String username) async {
  var url = await buildPath("Blog.API", "article", new List<String>());
  var data = jsonEncode({
    "Title": title,
    "Intro": "Provide a 128 character description of the article.",
    "Category": "Default",
    "ImageKey": '0`0',
    "Content": 'Content here.',
    "WrittenBy": username,
    "Public": false
  });

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("POST", url);
  request.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
  request.setRequestHeader(
      "Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

Future<HttpRequest> updateArticle(String key, String title, String intro,
    String category, String content, String imageKey, String username) async {
  var url = await buildPath("Blog.API", "article", new List<String>());
  final data = jsonEncode({
    "Key": key,
    "Body": {
      "Title": title,
      "Intro": intro,
      "Category": category,
      "ImageKey": imageKey,
      "Content": content,
      "WrittenBy": username,
      "Public": false
    }
  });

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("PUT", url);
  request.setRequestHeader(
      "Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

Future<HttpRequest> publishArticle(String key, String title, String intro,
    String category, String content, String imageKey, String username) async {
  var url = await buildPath("Blog.API", "article", new List<String>());
  final data = jsonEncode({
    "Key": key,
    "Body": {
      "Title": title,
      "Intro": intro,
      "Category": category,
      "ImageKey": imageKey,
      "Content": content,
      "WrittenBy": username,
      "Public": true
    }
  });

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("PUT", url);
  request.setRequestHeader(
      "Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

Future<HttpRequest> removeArticle(String key) async {
  var url = await buildPath("Blog.API", "article", [key]);

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("DELETE", url);
  request.setRequestHeader(
      "Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send();

  return compltr.future;
}

void onProgress(ProgressEvent e) {
  if (e.lengthComputable) {
    print('Progress... ${e.total}/${e.loaded}');
  }
}
