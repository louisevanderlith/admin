import 'dart:convert';
import 'dart:html';

import 'pathlookup.dart';
import 'profileform.dart';
import 'uploadapi.dart';

String _imageURL;
String _objKey;

void main() async {
  print('Profile Edit');
  var path = window.location.pathname;
  _objKey = path.substring(path.lastIndexOf('/') + 1);

  new ProfileForm('#frmBasicsite', _objKey, '#txtTitle', '#txtDescription',
      '#txtEmail', '#txtPhone', '#txtURL', '#uplProfileImg', '#btnSaveSite');

  _imageURL = await buildPath('Artifact.API', "upload", ["file"]);
  querySelectorAll('input[type="file"]').onChange.listen(uploadFile);
}

void uploadFile(Event e) {
  if (e.target is FileUploadInputElement) {
    FileUploadInputElement fileElem = e.target;
    var files = fileElem.files;

    var forAttr = fileElem.dataset['for'];
    var nameAttr = fileElem.dataset['name'];
    var idAttr = _objKey.replaceFirst('%60', '`');
    var ctrlID = fileElem.id;
    var infoObj = {"For": forAttr, "ItemName": nameAttr, "ItemKey": idAttr};

    if (files.length > 0) {
      File firstFile = files[0];

      doUpload(firstFile, infoObj, ctrlID);
    }
  }
}

void doUpload(File file, Map<String, String> infoObj, String ctrlID) {
  var formData = new FormData();
  formData.appendBlob("file", file);
  formData.append("info", jsonEncode(infoObj));

  createUpload(formData).then((obj) {
    print(obj.response);
    var resp = jsonDecode(obj.response);
    finishUpload(resp, infoObj, ctrlID);
  });
}

void finishUpload(dynamic obj, Map<String, String> infoObj, String ctrlID) {
  if (obj['Error'].length > 0) {
    print(obj['Error']);
    return;
  }

  var data = obj['Data'];
  var fullURL = "${_imageURL}/${data}";

  var imageHolder = querySelector("#${ctrlID.replaceFirst('Img', 'View')}");
  var uploader = querySelector("#${ctrlID}");

  imageHolder.classes.remove('is-hidden');
  imageHolder.setAttribute('src', fullURL);

  uploader.dataset['id'] = data;
  uploader.attributes.remove('required');
}
