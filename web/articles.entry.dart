import 'dart:convert';
import 'dart:html';

import 'package:mango_blog/blogapi.dart';
import 'package:mango_blog/bodies/article.dart';
import 'package:mango_ui/keys.dart';

void main() {
  querySelector('#btnAdd').onClick.listen(onAddClick);
  document.body.onClick.matches('.deleter').listen(onDeleteClick);
}

void onAddClick(MouseEvent e) async {
  final data = new Article('New Article', 'Short introduction', 'Default',
      new Key('0`0'), 'Content', 'System', false);
  final result = await createArticle(data);
  var obj = jsonDecode(result.response);

  if (result.status == 200) {
    final data = obj['Data'];
    final rec = data['Record'];
    final key = rec['K'];

    final redir = "/blog/articles/${key}";
    window.location.replace(redir);
  } else {
    print(obj['Error']);
  }
}

void onDeleteClick(MouseEvent e) async {
  final targt = e.matchingTarget;

  if (targt is ButtonElement) {
    final toDelete = targt.dataset["key"];
    final warn = "Are you sure you want to Delete ${toDelete}?";
    if (window.confirm(warn)) {
      final req = await deleteArticle(new Key(toDelete));
      final resp = jsonDecode(req.response);

      if (req.status == 200) {
        window.alert(resp["Data"]);
      } else {
        print(resp["Error"]);
      }
    }
  }
}
