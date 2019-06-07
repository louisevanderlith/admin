import 'dart:convert';
import 'dart:html';

import 'package:Admin.APP/blogapi.dart';

void main() {
  querySelector('#btnAdd').onClick.listen(onAddClick);
}

void onAddClick(MouseEvent e) async {
  final result = await createArticle('New Article', 'System');
  var obj = jsonDecode(result.response);

  if (result.status == 200) {
    final redir = "/create/${obj['Data'].Key}";
    window.location.replace(redir);
  } else {
    print(obj['Error']);
  }
}