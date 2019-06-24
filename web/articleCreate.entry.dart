import 'package:Admin.APP/articleform.dart';
import 'package:Admin.APP/keys.dart';

void main() {
  print('Article Create');
  new ArticleForm("#frmBlogCreate", getObjKey(), "#txtTitle", "#txtIntro", "#cboCategories", "#txtContent", "#uplHeaderImg", "#btnPreview",
      "#btnPublish", "#btnSave");
}
