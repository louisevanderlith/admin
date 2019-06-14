import 'package:Admin.APP/blogform.dart';

void main() {
  print('Blog Create');
  new BlogForm("#frmBlogCreate", "#txtTitle", "#txtContent", "#uplHeader", "#btnPreview",
      "#btnPublish", "#btnSave");
}
