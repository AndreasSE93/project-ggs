package packageManaging;

public class StageFlipper {
	public int packageID;
	public JoinMessage jm;
	
	public StageFlipper() {
		
	}
	
	public StageFlipper(JoinMessage jm) {
		this.packageID = jm.PacketID;
		this.jm = jm;
	}
	
	public StageFlipper(KickMessage km) {
		this.packageID = km.PacketID;
	}
}
