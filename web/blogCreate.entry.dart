import 'package:Admin.APP/blogform.dart';
import 'package:Admin.APP/keys.dart';

void main() {
  print('Blog Create');
  new BlogForm("#frmBlogCreate", getObjKey(), "#txtTitle", "#txtContent", "#uplHeader", "#btnPreview",
      "#btnPublish", "#btnSave");
}
