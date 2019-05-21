import 'package:Admin.APP/keys.dart';
import 'package:Admin.APP/models/profileform.dart';

void main() async {
  print('Profile Edit');

  new ProfileForm(
      '#frmBasicsite',
      getObjKey(),
      '#txtTitle',
      '#txtDescription',
      '#txtEmail',
      '#txtPhone',
      '#txtURL',
      '#uplProfileImg',
      '#frmHeader',
      '#btnAddHeader',
      '#frmPortfolio',
      '#btnAddPortfolio',
      '#frmSocialmedia',
      '#btnAddSocial',
      '#btnSave');
}
