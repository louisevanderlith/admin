import 'package:mango_ui/keys.dart';
import 'package:Admin.APP/profileform.dart';

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
      '#txtGTag',
      '#uplProfileImg',
      '#frmHeader',
      '#btnAddHeader',
      '#frmPortfolio',
      '#btnAddPortfolio',
      '#frmSocialmedia',
      '#btnAddSocial',
      '#btnSubmit');
}
