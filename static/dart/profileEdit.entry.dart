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
  querySelector('input[type="file"]').onChange.listen(uploadFile);
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
  //var reader = new FileReader();

  //reader.onLoadEnd.listen((ProgressEvent e) {
    //var formData = new Map<String, String>();
    var formData = new FormData();
    formData.appendBlob("file", file);
    formData.append("info", jsonEncode(infoObj));
    
    createUpload(formData).then((obj) => {finishUpload(obj, infoObj, ctrlID)});
  //});

  //reader.readAsArrayBuffer(file);
}

void finishUpload(
    Map<String, String> obj, Map<String, String> infoObj, String ctrlID) {
  var fullURL = _imageURL + "/" + obj["Data"];

  var imageHolder = querySelector("#${ctrlID.replaceFirst('Img', 'View')}");
  var uploader = querySelector("#${ctrlID}");

  imageHolder.attributes.remove('hidden');
  imageHolder.setAttribute('src', fullURL);

  uploader.dataset['id'] = obj['Data'];
  uploader.attributes.remove('required');
}
