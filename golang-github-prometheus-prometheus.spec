
%global debug_package   %{nil}

%global provider        github
%global provider_tld    com
%global project         prometheus
%global repo            prometheus
# https://github.com/prometheus/prometheus
%global provider_prefix %{provider}.%{provider_tld}/%{project}/%{repo}
%global import_path     %{provider_prefix}
# %commit is intended to be set by tito. The values in this spec file will not be kept up to date.
%{!?commit:
%global commit          bc6058c81272a8d938c05e75607371284236aadc
}
%global shortcommit     %(c=%{commit}; echo ${c:0:7})
%global gopathdir       %{_sourcedir}/go
%global upstream_ver    2.3.1
%global rpm_ver         %(v=%{upstream_ver}; echo ${v//-/_})
%global download_prefix %{provider}.%{provider_tld}/openshift/%{repo}

Name:		golang-%{provider}-%{project}-%{repo}
# Version and release information will be automatically managed by CD
# It will be kept in sync with OCP builds.
Version:	%{rpm_ver}
Release:	1.git%{shortcommit}%{?dist}
Summary:	The Prometheus monitoring system and time series database
License:	ASL 2.0
URL:		https://prometheus.io/
Source0:	https://%{download_prefix}/archive/%{commit}/%{repo}-%{commit}.tar.gz

# e.g. el6 has ppc64 arch without gcc-go, so EA tag is required
# promu-based packages FTBFS on aarch64 (#1487462)
ExclusiveArch:  %{?go_arches:%{go_arches}}%{!?go_arches:%{ix86} x86_64 %{arm} ppc64le s390x}
# If go_compiler is not set to 1, there is no virtual provide. Use golang instead.
BuildRequires: %{?go_compiler:compiler(go-compiler)}%{!?go_compiler:golang}
BuildRequires: prometheus-promu

%description
%{summary}

%package -n %{project}-prometheus
Summary:        %{summary}
Provides:       prometheus-prometheus = %{version}-%{release}
Obsoletes:      prometheus-prometheus < %{version}-%{release}

%description -n %{project}-prometheus
%{summary}

%prep
%setup -q -n %{repo}-%{commit}

%build
# Go expects a full path to the sources which is not included in the source
# tarball so create a link with the expected path
mkdir -p %{gopathdir}/src/%{provider}.%{provider_tld}/%{project}
GOSRCDIR=%{gopathdir}/src/%{import_path}
if [ ! -e "$GOSRCDIR" ]; then
  ln -s `pwd` "$GOSRCDIR"
fi
export GOPATH=%{gopathdir}

make build BUILD_PROMU=false

%install
install -d %{buildroot}%{_bindir}
install -D -p -m 0755 prometheus %{buildroot}%{_bindir}/prometheus
install -D -p -m 0755 promtool %{buildroot}%{_bindir}/promtool
install -d %{buildroot}%{_sysconfdir}/prometheus
install -D -p -m 0644 documentation/examples/prometheus.yml \
                      %{buildroot}%{_sysconfdir}/prometheus/prometheus.yml
install -D -p -m 0644 prometheus.service \
                      %{buildroot}%{_unitdir}/prometheus.service
install -D -p -m 0644 prometheus.sysconfig \
                      %{buildroot}%{_sysconfdir}/sysconfig/prometheus
install -d %{buildroot}%{_sharedstatedir}/prometheus
install -d %{buildroot}%{_datadir}/prometheus/console_libraries
install -D -p -m 0644 console_libraries/* %{buildroot}%{_datadir}/prometheus/console_libraries
install -d %{buildroot}%{_datadir}/prometheus/consoles
install -D -p -m 0644 consoles/* %{buildroot}%{_datadir}/prometheus/consoles

%files -n %{project}-prometheus
%license LICENSE NOTICE
%doc CHANGELOG.md CONTRIBUTING.md MAINTAINERS.md README.md
%{_bindir}/prometheus
%{_bindir}/promtool
%{_sysconfdir}/prometheus/prometheus.yml
%{_unitdir}/prometheus.service
%{_sysconfdir}/sysconfig/prometheus
%{_sharedstatedir}/prometheus
%{_datadir}/prometheus/console_libraries
%{_datadir}/prometheus/consoles

%changelog
* Mon Jun 25 2018 Simon Pasquier <spasquie@redhat.com> - 2.3.1-1
- Upgrade to 2.3.1

* Wed Mar 28 2018 Paul Gier <pgier@redhat.com> - 2.2.1-1
- Upgrade to 2.2.1

* Wed Feb 21 2018 Paul Gier <pgier@redhat.com> - 2.1.0-1
- Upgrade to 2.1.0

* Wed Jan 10 2018 Yaakov Selkowitz <yselkowi@redhat.com> - 2.0.0-2.git0a74f98
- Rebuilt for ppc64le, s390x enablement

* Wed Nov 08 2017 Paul Gier <pgier@redhat.com> - 2.0.0-1
- Upgrade to 2.0.0 ga

* Fri Oct 27 2017 Paul Gier <pgier@redhat.com> - 2.0.0_rc.2-1
- Upgrade to 2.0.0_rc.2

* Tue Oct 03 2017 Paul Gier <pgier@redhat.com> - 2.0.0_beta.5-1
- Upgrade to 2.0.0_beta.5 with openshift patches

* Fri Sep 01 2017 Paul Gier <pgier@redhat.com> - 2.0.0_beta.2-2.gitcb0f605
- Use sources from forked openshift project
- Don't build the duplicate binary package
- Strip debug info

* Wed Aug 23 2017 Paul Gier <pgier@redhat.com> - 2.0.0_beta.2-1.gita52f082
- First package for Openshift
