import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:Admin.APP/models/headeritem.dart';
import 'package:Admin.APP/models/portfolioitem.dart';
import 'package:Admin.APP/models/socialmediaitem.dart';

import '../pathlookup.dart';

Future<HttpRequest> createProfile(
    String title,
    String descr,
    String email,
    String phone,
    String url,
    String imageKey,
    List<PortfolioItem> portfolio,
    List<SocialmediaItem> socials,
    List<HeaderItem> headers) async {
  final data = jsonEncode({
    "Title": title,
    "Description": descr,
    "ContactEmail": email,
    "ContactPhone": phone,
    "URL": url,
    "ImageKey": imageKey,
    "PortfolioItems": portfolio,
    "SocialLinks": socials,
    "Headers": headers
  });

  return sendProfile("POST", data);
}

Future<HttpRequest> updateProfile(
    String key,
    String title,
    String descr,
    String email,
    String phone,
    String url,
    String imageKey,
    List<PortfolioItem> portfolio,
    List<SocialmediaItem> socials,
    List<HeaderItem> headers) async {
  final data = jsonEncode({
    "Key": key,
    "Body": {
      "Title": title,
      "Description": descr,
      "ContactEmail": email,
      "ContactPhone": phone,
      "URL": url,
      "ImageKey": imageKey,
      "PortfolioItems": portfolio,
      "SocialLinks": socials,
      "Headers": headers
    }
  });

  return sendProfile("PUT", data);
}

Future<HttpRequest> sendProfile(String method, String data) async {
  final url = await buildPath("Folio.API", "profile", new List<String>());

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open(method, url);
  request.setRequestHeader(
      "Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

void onProgress(ProgressEvent e) {
  if (e.lengthComputable) {
    print('Progress... ${e.total}/${e.loaded}');
  }
}
