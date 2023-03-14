Name:                forrealgo
Version:             1.0
Release:             1%{?dist}
Summary:             CLI For Real - CLI NLP tool for generating commands

License:             GPLv3
Source0:             %{name}-%{version}.tar.gz

BuildRequires:       golang
BuildRequires:       systemd-rpm-macros

Provides:            %{name} = %{version}

%description
CLI For Real - CLI NLP tool for generating commands

%global debug_package %{nil}

%prep
%autosetup


%build
go build -v -o %{name}


%install
install -Dpm 0755 %{name} %{buildroot}%{_bindir}/%{name}
install -Dpm 0755 config.json %{buildroot}%{_sysconfdir}/%{name}/config.json
install -Dpm 644 %{name}.service %{buildroot}%{_unitdir}/%{name}.service

%check
# go test should be here... :)

%post
%systemd_post %{name}.service

%preun
%systemd_preun %{name}.service

%files
%dir %{_sysconfdir}/%{name}
%{_bindir}/%{name}
%{_unitdir}/%{name}.service
%config(noreplace) %{_sysconfdir}/%{name}/config.json
