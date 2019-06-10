import 'dart:convert';
import 'dart:html';

import 'package:Admin.APP/services/blogapi.dart';

void main() {
  querySelector('#btnAdd').onClick.listen(onAddClick);
}

void onAddClick(MouseEvent e) async {
  final result = await createArticle('New Article', 'System');
  var obj = jsonDecode(result.response);

  if (result.status == 200) {
    final data = obj['Data'];
    final rec = data['Record'];
    final key = rec['K'];
    print('Key: ${key}');
    final redir = "/blog/${key}";
    window.location.replace(redir);
  } else {
    print(obj['Error']);
  }
}