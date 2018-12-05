// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TNewProxyServerForm struct {
    *vcl.TForm
    GroupBox1          *vcl.TGroupBox
    Label3             *vcl.TLabel
    LocalTypeCombox    *vcl.TComboBox
    Label4             *vcl.TLabel
    LocalIpEdit        *vcl.TEdit
    LocalPortEdit      *vcl.TEdit
    GroupBox2          *vcl.TGroupBox
    Label1             *vcl.TLabel
    SrcTypeComboBox    *vcl.TComboBox
    Label2             *vcl.TLabel
    SrcIpEdit          *vcl.TEdit
    SrcPortEdit        *vcl.TEdit
    GroupBox3          *vcl.TGroupBox
    TlsChk             *vcl.TCheckBox
    CrtEdit            *vcl.TEdit
    KeyEdit            *vcl.TEdit
    SelectCrtBtn       *vcl.TButton
    SelectKeyBtn       *vcl.TButton
    GroupBox4          *vcl.TGroupBox
    Label5             *vcl.TLabel
    AuthAddrEdit       *vcl.TEdit
    Label6             *vcl.TLabel
    AuthMemo           *vcl.TMemo
    Label7             *vcl.TLabel
    ConfigNameEdit     *vcl.TEdit
    OkButton           *vcl.TButton
    CancelButton       *vcl.TButton
    SelectCrtDialog    *vcl.TOpenDialog
    SelectKeyDialog    *vcl.TOpenDialog
    ActionList1        *vcl.TActionList
    SelectCrtAction    *vcl.TAction
    SelectKeyAction    *vcl.TAction
    EnableTlsAction    *vcl.TAction

    //::private::
    TNewProxyServerFormFields
}

var NewProxyServerForm *TNewProxyServerForm




// 以字节形式加载
// vcl.Application.CreateForm(newProxyServerFormBytes, &NewProxyServerForm)

var newProxyServerFormBytes = []byte("\x54\x50\x46\x30\x13\x54\x4E\x65\x77\x50\x72\x6F\x78\x79\x53\x65\x72\x76\x65\x72\x46\x6F\x72\x6D\x12\x4E\x65\x77\x50\x72\x6F\x78\x79\x53\x65\x72\x76\x65\x72\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x03\x45\x02\x06\x48\x65\x69\x67\x68\x74\x03\x56\x02\x03\x54\x6F\x70\x03\xEF\x00\x05\x57\x69\x64\x74\x68\x03\xDF\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x96\xB0\xE5\xA2\x9E\xE4\xBB\xA3\xE7\x90\x86\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x56\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xDF\x01\x0D\x44\x65\x73\x69\x67\x6E\x54\x69\x6D\x65\x50\x50\x49\x02\x78\x08\x4F\x6E\x43\x72\x65\x61\x74\x65\x07\x0A\x46\x6F\x72\x6D\x43\x72\x65\x61\x74\x65\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x07\x31\x2E\x38\x2E\x34\x2E\x30\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x31\x04\x4C\x65\x66\x74\x02\x04\x06\x48\x65\x69\x67\x68\x74\x02\x70\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xD7\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x9C\xAC\xE5\x9C\xB0\xE9\x85\x8D\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x57\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xD3\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xB1\xBB\xE5\x9E\x8B\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x0F\x4C\x6F\x63\x61\x6C\x54\x79\x70\x65\x43\x6F\x6D\x62\x6F\x78\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x7D\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x14\x09\x49\x74\x65\x6D\x49\x6E\x64\x65\x78\x02\x00\x0D\x49\x74\x65\x6D\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x04\x48\x54\x54\x50\x06\x05\x53\x4F\x43\x4B\x53\x06\x03\x53\x50\x53\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x04\x54\x65\x78\x74\x06\x04\x48\x54\x54\x50\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x34\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x02\x2F\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x49\x50\x3A\xE7\xAB\xAF\xE5\x8F\xA3\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0B\x4C\x6F\x63\x61\x6C\x49\x70\x45\x64\x69\x74\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x03\xA0\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x04\x54\x65\x78\x74\x06\x07\x30\x2E\x30\x2E\x30\x2E\x30\x00\x00\x05\x54\x45\x64\x69\x74\x0D\x4C\x6F\x63\x61\x6C\x50\x6F\x72\x74\x45\x64\x69\x74\x04\x4C\x65\x66\x74\x03\x00\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x02\x40\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x04\x54\x65\x78\x74\x06\x05\x33\x38\x30\x38\x30\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x32\x04\x4C\x65\x66\x74\x02\x04\x06\x48\x65\x69\x67\x68\x74\x02\x70\x03\x54\x6F\x70\x02\x70\x05\x57\x69\x64\x74\x68\x03\xD7\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE4\xB8\x8A\xE7\xBA\xA7\xE9\x85\x8D\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x57\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xD3\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xB1\xBB\xE5\x9E\x8B\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x0F\x53\x72\x63\x54\x79\x70\x65\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x7D\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x14\x09\x49\x74\x65\x6D\x49\x6E\x64\x65\x78\x02\x00\x0D\x49\x74\x65\x6D\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x04\x48\x54\x54\x50\x06\x05\x53\x4F\x43\x4B\x53\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x04\x54\x65\x78\x74\x06\x04\x48\x54\x54\x50\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x02\x2F\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x49\x50\x3A\xE7\xAB\xAF\xE5\x8F\xA3\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x09\x53\x72\x63\x49\x70\x45\x64\x69\x74\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x03\xA0\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x04\x54\x65\x78\x74\x06\x0F\x32\x35\x35\x2E\x32\x35\x35\x2E\x32\x35\x35\x2E\x32\x35\x35\x00\x00\x05\x54\x45\x64\x69\x74\x0B\x53\x72\x63\x50\x6F\x72\x74\x45\x64\x69\x74\x04\x4C\x65\x66\x74\x03\x00\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x02\x40\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x04\x54\x65\x78\x74\x06\x05\x33\x38\x30\x38\x30\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x33\x04\x4C\x65\x66\x74\x02\x04\x06\x48\x65\x69\x67\x68\x74\x03\x88\x00\x03\x54\x6F\x70\x03\xE0\x00\x05\x57\x69\x64\x74\x68\x03\xD7\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE5\x8A\xA0\xE5\xAF\x86\xE9\x85\x8D\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x6F\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xD3\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x06\x54\x6C\x73\x43\x68\x6B\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x18\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x02\x54\x06\x41\x63\x74\x69\x6F\x6E\x07\x0F\x45\x6E\x61\x62\x6C\x65\x54\x6C\x73\x41\x63\x74\x69\x6F\x6E\x07\x43\x68\x65\x63\x6B\x65\x64\x09\x05\x53\x74\x61\x74\x65\x07\x09\x63\x62\x43\x68\x65\x63\x6B\x65\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x05\x54\x45\x64\x69\x74\x07\x43\x72\x74\x45\x64\x69\x74\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x03\x38\x01\x05\x43\x6F\x6C\x6F\x72\x07\x08\x63\x6C\x53\x69\x6C\x76\x65\x72\x08\x52\x65\x61\x64\x4F\x6E\x6C\x79\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x05\x54\x45\x64\x69\x74\x07\x4B\x65\x79\x45\x64\x69\x74\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x03\x38\x01\x05\x43\x6F\x6C\x6F\x72\x07\x08\x63\x6C\x53\x69\x6C\x76\x65\x72\x08\x52\x65\x61\x64\x4F\x6E\x6C\x79\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0C\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x42\x74\x6E\x04\x4C\x65\x66\x74\x03\x50\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x02\x5E\x06\x41\x63\x74\x69\x6F\x6E\x07\x0F\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x41\x63\x74\x69\x6F\x6E\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0C\x53\x65\x6C\x65\x63\x74\x4B\x65\x79\x42\x74\x6E\x04\x4C\x65\x66\x74\x03\x50\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x02\x5E\x06\x41\x63\x74\x69\x6F\x6E\x07\x0F\x53\x65\x6C\x65\x63\x74\x4B\x65\x79\x41\x63\x74\x69\x6F\x6E\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x34\x04\x4C\x65\x66\x74\x02\x04\x06\x48\x65\x69\x67\x68\x74\x03\x98\x00\x03\x54\x6F\x70\x03\x68\x01\x05\x57\x69\x64\x74\x68\x03\xD7\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xAE\xA4\xE8\xAF\x81\xE9\x85\x8D\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x7F\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xD3\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x35\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x07\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xAE\xA4\xE8\xAF\x81\xE5\x9C\xB0\xE5\x9D\x80\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0C\x41\x75\x74\x68\x41\x64\x64\x72\x45\x64\x69\x74\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x07\x05\x57\x69\x64\x74\x68\x03\xEC\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x36\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE7\x94\xA8\xE6\x88\xB7\xE5\xAF\x86\xE7\xA0\x81\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x4D\x65\x6D\x6F\x08\x41\x75\x74\x68\x4D\x65\x6D\x6F\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x48\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x03\xEC\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x37\x04\x4C\x65\x66\x74\x02\x10\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x03\x08\x02\x05\x57\x69\x64\x74\x68\x02\x5A\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\xE9\x85\x8D\xE7\xBD\xAE\xE6\x96\x87\xE4\xBB\xB6\xE5\x90\x8D\xE7\xA7\xB0\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0E\x43\x6F\x6E\x66\x69\x67\x4E\x61\x6D\x65\x45\x64\x69\x74\x04\x4C\x65\x66\x74\x03\x80\x00\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x03\x08\x02\x05\x57\x69\x64\x74\x68\x03\xE4\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x08\x4F\x6B\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x02\x58\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x03\x30\x02\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xA1\xAE\xE5\xAE\x9A\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0D\x4F\x6B\x42\x75\x74\x74\x6F\x6E\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0C\x43\x61\x6E\x63\x65\x6C\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x03\x06\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x03\x30\x02\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8F\x96\xE6\xB6\x88\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x11\x43\x61\x6E\x63\x65\x6C\x42\x75\x74\x74\x6F\x6E\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x06\x00\x00\x0B\x54\x4F\x70\x65\x6E\x44\x69\x61\x6C\x6F\x67\x0F\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x44\x69\x61\x6C\x6F\x67\x04\x6C\x65\x66\x74\x02\x68\x03\x74\x6F\x70\x03\xE8\x00\x00\x00\x0B\x54\x4F\x70\x65\x6E\x44\x69\x61\x6C\x6F\x67\x0F\x53\x65\x6C\x65\x63\x74\x4B\x65\x79\x44\x69\x61\x6C\x6F\x67\x04\x6C\x65\x66\x74\x03\xE0\x00\x03\x74\x6F\x70\x03\xE8\x00\x00\x00\x0B\x54\x41\x63\x74\x69\x6F\x6E\x4C\x69\x73\x74\x0B\x41\x63\x74\x69\x6F\x6E\x4C\x69\x73\x74\x31\x04\x6C\x65\x66\x74\x03\x67\x01\x03\x74\x6F\x70\x03\x88\x01\x00\x07\x54\x41\x63\x74\x69\x6F\x6E\x0F\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x41\x63\x74\x69\x6F\x6E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE9\x80\x89\xE6\x8B\xA9\xE8\xAF\x81\xE4\xB9\xA6\x09\x4F\x6E\x45\x78\x65\x63\x75\x74\x65\x07\x16\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x41\x63\x74\x69\x6F\x6E\x45\x78\x65\x63\x75\x74\x65\x00\x00\x07\x54\x41\x63\x74\x69\x6F\x6E\x0F\x53\x65\x6C\x65\x63\x74\x4B\x65\x79\x41\x63\x74\x69\x6F\x6E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE9\x80\x89\xE6\x8B\xA9\x4B\x45\x59\x09\x4F\x6E\x45\x78\x65\x63\x75\x74\x65\x07\x16\x53\x65\x6C\x65\x63\x74\x4B\x65\x79\x41\x63\x74\x69\x6F\x6E\x45\x78\x65\x63\x75\x74\x65\x00\x00\x07\x54\x41\x63\x74\x69\x6F\x6E\x0F\x45\x6E\x61\x62\x6C\x65\x54\x6C\x73\x41\x63\x74\x69\x6F\x6E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE5\x90\xAF\xE7\x94\xA8\xE5\x8A\xA0\xE5\xAF\x86\x09\x4F\x6E\x45\x78\x65\x63\x75\x74\x65\x07\x16\x45\x6E\x61\x62\x6C\x65\x54\x6C\x73\x41\x63\x74\x69\x6F\x6E\x45\x78\x65\x63\x75\x74\x65\x00\x00\x00\x00")
