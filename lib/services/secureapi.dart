import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:Admin.APP/models/roleitem.dart';
import 'package:Admin.APP/pathlookup.dart';

Future<HttpRequest> updateRoles(String key, List<RoleItem> items) async {
  var url = await buildPath("Secure.API", "user", [key]);
  final data = jsonEncode(items);

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("PUT", url);
  request.setRequestHeader("Authorization", "Bearer " + window.localStorage['avosession']);
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
