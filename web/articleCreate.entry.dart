import 'package:Admin.APP/articleform.dart';
import 'package:mango_ui/keys.dart';

void main() {
  print('Article Create');
  new ArticleForm("#frmBlogCreate", getObjKey(), "#txtTitle", "#txtIntro", "#cboCategories", "#txtContent", "#uplHeaderImg", "#hdnAuthor", "#chkPublic", "#btnPreview",
      "#btnPublish", "#btnSave");
}
