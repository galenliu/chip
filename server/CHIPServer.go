package server

import (
	"github.com/galenliu/chip/access"
	"github.com/galenliu/chip/config"
	"github.com/galenliu/chip/controller"
	"github.com/galenliu/chip/credentials"
	"github.com/galenliu/chip/inet/udp_endpoint"
	"github.com/galenliu/chip/lib"
	"github.com/galenliu/chip/messageing"
	sd "github.com/galenliu/chip/server/dnssd"
	"github.com/galenliu/chip/storage"
	"github.com/galenliu/chip/transport"
	log "github.com/sirupsen/logrus"
	"net"
)

type AppDelegate interface {
	OnCommissioningSessionStarted()
	OnCommissioningSessionStopped()
	OnCommissioningWindowOpened()
	OnCommissioningWindowClosed()
}

type Server struct {
	mSecuredServicePort            uint16
	mUnsecuredServicePort          uint16
	mOperationalServicePort        uint16
	mUserDirectedCommissioningPort uint16
	mInterfaceId                   net.Interface
	mDnssd                         sd.DnssdServer
	mFabrics                       *credentials.FabricTable
	mCommissioningWindowManager    *sd.CommissioningWindowManager
	mDeviceStorage                 storage.PersistentStorageDelegate //unknown
	mAccessControl                 access.AccessControler
	mSessionResumptionStorage      any
	mExchangeMgr                   messageing.ExchangeManager
	mAttributePersister            lib.AttributePersistenceProvider //unknown
	mAclStorage                    *AclStorage
	mTransports                    transport.TransportManager
	mListener                      any
}

func NewServer(initParams *CommonCaseDeviceServerInitParams) *Server {
	s := &Server{}
	log.Printf("app server initializing")

	var err error
	s.mUnsecuredServicePort = initParams.OperationalServicePort
	s.mSecuredServicePort = initParams.UserDirectedCommissioningPort
	s.mInterfaceId = initParams.InterfaceId

	s.mCommissioningWindowManager.SetAppDelegate(initParams.AppDelegate)

	s.mDnssd = sd.NewServer()
	s.mDnssd.SetFabricTable(s.mFabrics)
	s.mCommissioningWindowManager = sd.CommissioningWindowManager{}.Init(&s)
	//s.mCommissioningWindowManager.SetAppDelegate(initParams.AppDelegate)

	// Initialize KvsPersistentStorageDelegate-based storage
	s.mDeviceStorage = initParams.PersistentStorageDelegate
	s.mSessionResumptionStorage = initParams.SessionResumptionStorage

	// Set up attribute persistence before we try to bring up the data model
	// handler.
	if s.mAttributePersister != nil {
		err = s.mAttributePersister.Init(s.mDeviceStorage)
		if err != nil {
			log.Panic(err.Error())
		}
	}

	if s.mFabrics != nil {
		err = s.mFabrics.Init(s.mDeviceStorage)
		if err != nil {
			log.Panic(err.Error())
		}
	}

	//少sDeviceTypeResolver参数
	if s.mAccessControl != nil {
		err = s.mAccessControl.Init(initParams.AccessDelegate)
		if err != nil {
			log.Panic(err.Error())
		}
	}

	s.mDnssd.SetFabricTable(s.mFabrics)
	s.mDnssd.SetCommissioningModeProvider(s.mCommissioningWindowManager)

	//mGroupsProvider = initParams.groupDataProvider;
	//SetGroupDataProvider(mGroupsProvider);
	//
	//deviceInfoprovider = DeviceLayer::GetDeviceInfoProvider();
	//if (deviceInfoprovider)
	//{
	//	deviceInfoprovider->SetStorageDelegate(mDeviceStorage);
	//}

	// This initializes clusters, so should come after lower level initialization.
	//不知道干什么的
	controller.InitDataModelHandler(s.mExchangeMgr)

	params := transport.UdpListenParameters{}
	params.SetListenPort(s.mOperationalServicePort)
	params.SetNativeParams(initParams.EndpointNativeParams)
	s.mTransports, err = transport.NewUdpTransport(udp_endpoint.UDPEndpoint{}, params)

	//s.mListener, err = mdns.IntGroupDataProviderListener(s.mTransports)
	if err != nil {
		log.Panic(err.Error())
	}

	//dnssd.ResolverInstance().Init(udp_endpoint.UDPEndpoint{})

	s.mDnssd.SetSecuredPort(s.mOperationalServicePort)
	s.mDnssd.SetUnsecuredPort(s.mUserDirectedCommissioningPort)
	s.mDnssd.SetInterfaceId(s.mInterfaceId)

	if s.GetFabricTable() != nil {
		if s.GetFabricTable().FabricCount() != 0 {
			if config.ChipConfigNetworkLayerBle {
				//TODO
				//如果Fabric不为零，那么设备已经被添加
				//可以在这里关闭蓝牙
			}
		}
	}

	//如果设备开启了自动配对模式，进入模式
	if config.ChipDeviceConfigEnablePairingAutostart {
		s.GetFabricTable().DeleteAllFabrics()
		err = s.mCommissioningWindowManager.OpenBasicCommissioningWindow()
		if err != nil {
			log.Panic(err.Error())
		}
	}
	s.mDnssd.StartServer()

	return s
}

// GetFabricTable 返回CHIP服务中的Fabric
func (s Server) GetFabricTable() *credentials.FabricTable {
	return s.mFabrics
}

func (s Server) Shutdown() {

}

func (s *Server) StartServer() error {
	return nil
}