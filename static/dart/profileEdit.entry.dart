import 'dart:convert';
import 'dart:html';

import 'pathlookup.dart';
import 'profileform.dart';
import 'uploadapi.dart';

String _imageURL;

void main() async {
  print('Profile Edit');
  var path = window.location.pathname;
  var currentID = path.substring(path.lastIndexOf('/') + 1);

  new ProfileForm('#frmBasicSite', currentID, '#txtTitle', '#txtDescription',
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
    var idAttr = fileElem.dataset['itemid'];
    var ctrlID = fileElem.id;
    var infoObj = {"For": forAttr, "ItemName": nameAttr, "ItemID": idAttr};

    if (files.length > 0) {
      var formData = new FormData();
      File firstFile = files[0];
      formData.append('file', firstFile.toString());
      formData.append('info', jsonEncode(infoObj));

      doUpload(formData, infoObj, ctrlID);
    }
  }
}

void doUpload(FormData formData, Map<String, String> infoObj, String ctrlID) {
  var success = (obj) => {finishUpload(obj, infoObj, ctrlID)};

  createUpload(formData, success);
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
